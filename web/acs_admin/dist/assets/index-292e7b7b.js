import{a as E,b as B,c as S}from"./index-77a01a79.js";import{l as N}from"./lodash-fa97179f.js";import{_ as T,r as v,o as U,b as u,c as H,e as f,i as M,f as y,n as w,h as e,l as F,w as l,x as c,F as R,E as g,a as j,m as G,g as x}from"./index-07fbba9f.js";import"./request-675639ab.js";const z={__name:"formSetting",setup(D){const i=v(!1),p=v(!1),a=v({}),h=v(!1);U(()=>{k()});const b=()=>{p.value=!0;let m=N.cloneDeep(a.value);B(m).then(()=>{p.value=!1,g.success("保存成功")})},k=async()=>{let m=await E(),o=JSON.parse(m.data);a.value=o,i.value=!0,m.is_save==="1"&&(h.value=!0)},I=()=>{p.value=!0;let m=N.cloneDeep(a.value);S(m).then(()=>{p.value=!1,g.success("测试连接成功")}).catch(()=>{p.value=!1})};return(m,o)=>{const V=u("el-alert"),s=u("el-divider"),C=u("el-switch"),_=u("el-form-item"),n=u("el-col"),r=u("el-row"),d=u("el-input"),J=u("el-option"),K=u("el-select"),$=u("el-button"),q=u("el-form"),P=H("loading");return f(),M(R,null,[h.value?w("",!0):(f(),y(V,{key:0,title:"当前配置文件未保存到数据库",type:"info",style:{"margin-bottom":"10px"},"show-icon":"",closable:!1})),e(V,{title:"注意：此处修改后会覆盖默认配置，删除后会恢复默认配置。修改配置后，需重启服务生效。",type:"info","show-icon":"",closable:!1}),i.value?F((f(),y(q,{key:1,"label-width":"150px",class:"form_class",size:"large"},{default:l(()=>[e(s,{"content-position":"left"},{default:l(()=>[c("基本配置")]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"调试模式"},{default:l(()=>[e(C,{"model-value":a.value.debug,"onUpdate:modelValue":o[0]||(o[0]=t=>a.value.debug=t),"inline-prompt":!0},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"后台访问监控"},{default:l(()=>[e(C,{"model-value":a.value.open_admin_record,"onUpdate:modelValue":o[1]||(o[1]=t=>a.value.open_admin_record=t),"inline-prompt":!0},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"HTTP端口"},{default:l(()=>[e(d,{"model-value":a.value.http_serve_port,"onUpdate:modelValue":o[2]||(o[2]=t=>a.value.http_serve_port=t),type:"number",placeholder:"http运行端口"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"GRPC端口"},{default:l(()=>[e(d,{"model-value":a.value.grpc_serve_port,"onUpdate:modelValue":o[3]||(o[3]=t=>a.value.grpc_serve_port=t),type:"number",placeholder:"GRPC运行端口"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"记录留存"},{default:l(()=>[e(d,{"model-value":a.value.record_save_days,"onUpdate:modelValue":o[4]||(o[4]=t=>a.value.record_save_days=t),type:"number",placeholder:"记录留存天数"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12})]),_:1}),e(s,{"content-position":"left"},{default:l(()=>[c("登陆信息")]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"后台登陆用户名"},{default:l(()=>[e(d,{"model-value":a.value.admin_user,"onUpdate:modelValue":o[5]||(o[5]=t=>a.value.admin_user=t),type:"text",placeholder:"后台登陆用户名"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"后台登陆密码"},{default:l(()=>[e(d,{"model-value":a.value.admin_password,"onUpdate:modelValue":o[6]||(o[6]=t=>a.value.admin_password=t),"show-password":"",type:"password",placeholder:"后台登陆密码"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"HTTP User"},{default:l(()=>[e(d,{"model-value":a.value.http_user,"onUpdate:modelValue":o[7]||(o[7]=t=>a.value.http_user=t),type:"text",placeholder:"http服务用户"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"HTTP Pass"},{default:l(()=>[e(d,{"model-value":a.value.http_password,"onUpdate:modelValue":o[8]||(o[8]=t=>a.value.http_password=t),"show-password":"",type:"password",placeholder:"http服务密码"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"GRPC User"},{default:l(()=>[e(d,{"model-value":a.value.grpc_user,"onUpdate:modelValue":o[9]||(o[9]=t=>a.value.grpc_user=t),type:"text",placeholder:"grpc服务用户"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"GRPC Pass"},{default:l(()=>[e(d,{"model-value":a.value.grpc_password,"onUpdate:modelValue":o[10]||(o[10]=t=>a.value.grpc_password=t),"show-password":"",type:"password",placeholder:"grpc服务密码"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"Auth SECRET"},{default:l(()=>[e(d,{"model-value":a.value.secret_key,"onUpdate:modelValue":o[11]||(o[11]=t=>a.value.secret_key=t),type:"text",placeholder:"登陆SECRET"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"Server Expire Time"},{default:l(()=>[e(d,{"model-value":a.value.jwt_expire_time,"onUpdate:modelValue":o[12]||(o[12]=t=>a.value.jwt_expire_time=t),type:"number",placeholder:"服务登陆过期时间"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"Admin Expire Time"},{default:l(()=>[e(d,{"model-value":a.value.jwt_admin_expire_time,"onUpdate:modelValue":o[13]||(o[13]=t=>a.value.jwt_admin_expire_time=t),type:"number",placeholder:"后台管理登陆过期时间"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12})]),_:1}),e(s,{"content-position":"left"},{default:l(()=>[c("数据入库策略")]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"单次导入条数"},{default:l(()=>[e(d,{"model-value":a.value.db_batch_insert_num,"onUpdate:modelValue":o[14]||(o[14]=t=>a.value.db_batch_insert_num=t),type:"number",placeholder:"数据库单次插入数据条数"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"自动保存间隔"},{default:l(()=>[e(d,{"model-value":a.value.db_batch_insert_gap,"onUpdate:modelValue":o[15]||(o[15]=t=>a.value.db_batch_insert_gap=t),type:"number",placeholder:"数据每隔多少秒保存"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(s,{"content-position":"left"},{default:l(()=>[c("日志策略")]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"日志保存路径"},{default:l(()=>[e(d,{"model-value":a.value.log_file_path,"onUpdate:modelValue":o[16]||(o[16]=t=>a.value.log_file_path=t),type:"text",placeholder:"日志保存路径"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12},{default:l(()=>[e(_,{label:"日志保存天数"},{default:l(()=>[e(d,{"model-value":a.value.log_file_max_age,"onUpdate:modelValue":o[17]||(o[17]=t=>a.value.log_file_max_age=t),type:"number",placeholder:"日志文件最大保存天数"},null,8,["model-value"])]),_:1})]),_:1})]),_:1}),e(r,null,{default:l(()=>[e(n,{span:12},{default:l(()=>[e(_,{label:"日志切割间隔(小时)"},{default:l(()=>[e(d,{"model-value":a.value.log_file_gap,"onUpdate:modelValue":o[18]||(o[18]=t=>a.value.log_file_gap=t),type:"text",placeholder:"日志切割间隔(小时)"},null,8,["model-value"])]),_:1})]),_:1}),e(n,{span:12})]),_:1}),e(s,{"content-position":"left"},{default:l(()=>[c("数据库连接信息")]),_:1}),e(V,{style:{margin:"10px"},title:"注意：由于读取配置信息需要先连接数据库，所以数据库连接信息需要在配置文件修改。",type:"info","show-icon":"",closable:!1}),w("",!0),e(s),e(_,null,{default:l(()=>[e($,{type:"primary",onClick:b,plain:""},{default:l(()=>[c("保存")]),_:1})]),_:1})]),_:1})),[[P,p.value]]):w("",!0)],64)}}},A=T(z,[["__scopeId","data-v-d3f5c31b"]]),L=x("div",{class:"card-header"},[x("span",null,"系统设置")],-1),O={class:"config_item"},le={__name:"index",setup(D){const i=j(G()),p=[{path:"/",name:"首页"},{path:"",name:"系统设置"}];return U(()=>{i.SetTitle("系统设置"),i.SetBreadcrumb(p)}),(a,h)=>{const b=u("el-card");return f(),y(b,{shadow:"never"},{header:l(()=>[L]),default:l(()=>[x("div",O,[e(A)])]),_:1})}}};export{le as default};
