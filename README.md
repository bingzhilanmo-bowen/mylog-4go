# mylog-4go
- Use Logrus
- Set Default Config

# Go语言项目如何接入

## 接入

### 配置说明
|  配置参数   | 含义  |  类型  |  默认值  |
|  ----  | ----  | ----  | ----  |
| LogPath  | 日志文件落盘的位置 | string | /var/logs/nlog |
| Stdout  | 是否开启控制台输出 | boolean | false |
| LogFileMaxSize  | 单个日志文件的文件大小 | int | 200M |
| LogFileMaxAge  | 日志文件保留最近多少天 | int | 7 |
| LocalTime  | 是否使用本地时间做日志切割 | boolean | true |
| Compress  | 是否开启日志压缩 | boolean | false |
| OpenPrometheus  | 是否开启Prometheus Metric功能 | boolean | false |
| UseThreadLocal  | 是否支持ThreadLocal方式获取日志的公有Filed | boolean | false |


### 项目接入

```$xslt
import (
log "github.com/bingzhilanmo-bowen/mylog-4go"
）

// 初始化配置 也可以不设置有默认配置
func nlogConfig() {

	log.ConfigLog(&log.LoggerConfig{
		OpenPrometheus:true,
	})

}


func SimpleMethod(){
   // user log
   log.Infof("mylog log example %s ", " Im Data !!!")
   log.Debugf("mylog log debug example %s ", " Im Data Debug !!!")
   
   // with context 这类方法会去取traceId和requestId
   log.InfofCtx(ctx, "mylog log with ctx example %s ", " Im Data !!!")
   
   // metric log 
   log.Count(nlog.NewMetricBuilder()
                 .Metric("bowenD")
                 .Tags("t1","v1")
                 .TagsOnFly("tf1", "vfr1").Build())
                 
   // business log
   log.BusinessLog(&BisTest{Name: "HydxinDefault", Age: 39, Six: 1})    
   
   
   auditLog := log.NewBuilderAudit().Operator("bowen", "email", "xxxx").
		Action("INSERT").OptObject("ORG INFO").Tags("k", "v").RawData(RAW).UpdateData(UP).Build()

   log.Audit(auditLog)          
   
}
```
**Web 项目需要自行添加metrics的 Api暴露**

