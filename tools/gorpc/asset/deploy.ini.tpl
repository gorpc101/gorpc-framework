;----------------------织云包信息--------------
[pkg]
;Warning: This deploy ability musn't be used for online environment!
;http://yun.isd.com/index.php/package/?product=other&package=${zhiyunPackageName}
;织云包url中对应的product,eg:other
product=devtest

;织云包url中对应的package,eg:ilive_link_mic_test
package={{.ServerName}}_test
version=1.0.0

;----------------------编译文件信息--------------
[build]
;idc机器，且能够访问以下nfs. 可不修改
ip=10.100.73.163

;nfs文件夹. 可不修改
nfs=/nfs/now_ci/pkg/

;要上传的文件，多个文件用英文逗号分隔，eg:libilive_link_mic_svr.so,conf，该目录与deployPath一一对应
sourceFile={{.ServerName}},run.sh,../conf/service.ini,../conf/monitor.ini,../conf/trace.ini,../conf/log.ini

;部署到zhiyun包的相对路径，多个目录用英文逗号分隔，该目录与sourceFile一一对应，eg:conf/,lib/
deployPath=bin/,bin/,conf/,conf/,conf/,conf/

;本次包发布者，填自己rtx,需要在~/.goneat/deployspec.json里配置
operator={{.Author}}

;本次版本的备注信息，如r2154:更新so
remark=服务优化

;----------------------部署信息-----------------
[deploy]
;Warning: 请修改测试机IP，及是否自动部署,ip需要在~/.goneat/deployspec.json里配置
;enable,是否部署，当ips不为空且install=true时会部署，该参数可通过--install控制
install=true
;要部署的机器，多个ip用英文逗号分隔
ips={{.Ips}}
