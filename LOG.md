Cara untuk set up log ke aplikasi berbeda-beda tergantung struktur aplikasi dan tools log yang digunakan. Di artikel saya disini: https://medium.com/@halosyahrul/logging-14-best-practice-yang-seorang-developer-harus-tau-d8eddac25487, saya menyebutkan untuk menhindari vendor lock-in, approach menggunakan wrapper. Ini opsi pertama yang akan saya coba:

Di branch logging, saya membuat satu buah package baru bernama `log`. Di dalam package ini saya buat sebuah file `log.go`, isinya interface `Logger`, interface ini berisi method-method yang berfunsi untuk menulis log dengan berbagai level dan parameter, dan deklarasi sebuah variabel `Log` yang bertipe interface `Logger`. Dan terakhir sebuah function `InitLogger()` untuk menginisialisasi implementasi logger.

```go
package log

var Log Logger

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func InitLogger() {
	// Log = NewLogrusLogger() -- construct impl logrus dilakukan disini
}
```

Kemudian, setelah interfacenya sudah jadi, lalu siap lanjut ke implementasinya. Sebagai implementasi, saya akan menggunakan [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus). Oke saya tinggal membuat file baru untuk  implementasinya, file ini saya beri nama `logrus.go`, masih bertempat di package `log`.

Terlebih dulu saya buat implementasi untuk interface `Logger` dengan membungkus `*logrus.Logger` di dalam sebuah struct. Kemudian disemua method implementasinya bisa diisi dengan menggunakan method logging dari logrus.

```go
type LoggerWrapper struct {
	LogrusLogger *logrus.Logger
}

func (logger *LoggerWrapper) Info(args ...interface{}) {
	logger.LogrusLogger.Info(args)
}

func (logger *LoggerWrapper) Debug(args ...interface{}) {
	logger.LogrusLogger.Debug(args)
}

func (logger *LoggerWrapper) Errorf(format string, args ...interface{}) {
	logger.LogrusLogger.Errorf(format, args)
}

func (logger *LoggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.LogrusLogger.Fatalf(format, args)
}

func (logger *LoggerWrapper) Fatal(args ...interface{}) {
	logger.LogrusLogger.Fatal(args)
}

func (logger *LoggerWrapper) Infof(format string, args ...interface{}) {
	logger.LogrusLogger.Infof(format, args)
}

func (logger *LoggerWrapper) Warnf(format string, args ...interface{}) {
	logger.LogrusLogger.Warnf(format, args)
}

func (logger *LoggerWrapper) Debugf(format string, args ...interface{}) {
	logger.LogrusLogger.Debugf(format, args)
}

func (logger *LoggerWrapper) Printf(format string, args ...interface{}) {
	logger.LogrusLogger.Infof(format, args)
}

func (logger *LoggerWrapper) Println(args ...interface{}) {
	logger.LogrusLogger.Info(args, "\n")
}
```

Kemudian, sebagai constructor saya buat fungsi `NewLogrusLogger()` dengan return value interface `Logger`. Di function ini, saya create new instance logrus. Lalu saya tambahkan beberapa config untuk set level, source, dan formatter log. Masih banyak configurasi lain yang bisa dieksplor di logrus.

```go
func NewLogrusLogger() Logger {
	//  pembuatan instance baru untuk logrus
	logrusLogger := logrus.New()
	//  pengaturan logging level, disini saya menggunakan level Trace
	//  sehingga terlihat semua log untuk debugging
	logrusLogger.SetLevel(logrus.TraceLevel)
	//  ini untuk menampilkan log source (filename dan filenumber)
	logrusLogger.SetReportCaller(true)
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		//  untuk menentukan format timestamp
		FullTimestamp: true,
		//  untuk mengatur warna log di console, saya matikan karena saya lebih suka log saya berwarna.
		DisableColors: false,
		//  ini merupakan func untuk menulis log dengan lebih rapi.
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	return &LoggerWrapper{LogrusLogger: logrusLogger}
}
```

Setelah semua konfigurasi wrapper dan impl-nya selesai. Terakhir adalah menginit log di file `main.go` agar instance logrus dan semua confignya terload.