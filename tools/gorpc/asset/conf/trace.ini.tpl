[zipkin]
enabled = false                                             #是否启用zipkin trace
service.name = {{.ServerName}}                              #当前服务名称(span endpoint)
service.addr = *:8000                                       #当前服务地址(span endpoint)
collector.addr = http://9.24.146.130:8080/api/v1/spans      #zipkin collector接口地址(env.online)
collector.addr.dev = http://10.49.135.152:8080/api/v1/spans #zipkin collector接口地址(env.test)
traceId128bits = true                                       #是否启用128bits traceId

[jaeger]
enabled = false                                             #是否启用jaeger trace(暂未验证兼容性)

[tianjige]
enabled = false                                             #是否启用天机阁 trace
service.name = {{.ServerName}}                              #当前服务名称(span endpoint)
service.addr = *:8000                                       #当前服务地址(span endpoint)
collector.addr = 10.101.192.39:9092                         #天机阁 collector接口地址(env.online)
collector.addr.dev = 10.101.192.79:9092                     #天机阁 collector接口地址(env.test)
traceId128bits = true                                       #是否启用128bits traceId
appid = ContentPlatCenter                                   #天机阁 申请的业务id
