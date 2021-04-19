module tim-go

go 1.14

require (
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/casbin/gorm-adapter v1.0.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.12
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/spf13/viper v1.6.2
)

replace github.com/casbin/gorm-adapter => ./gorm-adapterv1.0.0
