import{r as j}from"./request-675639ab.js";import{r as F}from"./index-89cf8bcc.js";import{_ as J,a as K,m as Q,r as _,o as W,d as b,z as X,b as r,c as Z,e as p,f as D,w as n,g as w,h as e,i as T,j as E,F as I,u as V,A as ee,x as le,k as ae,l as U,q as te,s as oe}from"./index-07fbba9f.js";import{T as ne}from"./time_utils-1ad82d01.js";const ue=u=>j({url:`/record/list?page=${u.page}&size=${u.size}&start_time=${u.start_time}&end_time=${u.end_time}&group_id=${u.group_id}&url=${u.url}&method=${u.method}&ip=${u.ip}`,method:"get",data:u});const de=u=>(te("data-v-c535c589"),u=u(),oe(),u),re=de(()=>w("div",{class:"card-header"},[w("span",null,"请求记录列表")],-1)),se={class:"my_table_header"},ie={class:"page_lk"},_e={__name:"index",setup(u){const Y=K(Q()),v=_(!1),y=_([]),z=[{path:"/",name:"首页"},{path:"",name:"请求记录列表"}],a=_({group_id:"",url:"",method:"",ip:"",start_time:"",end_time:""}),$=[{val:"",name:"全部"},{val:"GET",name:"GET"},{val:"POST",name:"POST"},{val:"PUT",name:"PUT"},{val:"DELETE",name:"DELETE"},{val:"PATCH",name:"PATCH"},{val:"HEAD",name:"HEAD"},{val:"OPTIONS",name:"OPTIONS"}],f=_([{id:"",name:"全部"}]),A=ne(),P=_([]),m=_(1),S=_(0),x=_(10);W(async()=>{Y.SetTitle("请求记录列表"),Y.SetBreadcrumb(z),await G(),await h()});const h=async()=>{v.value=!0;let s={page:m.value,size:x.value,start_time:a.value.start_time,end_time:a.value.end_time,group_id:a.value.group_id,url:a.value.url,method:a.value.method,ip:a.value.ip},l=await ue(s);v.value=!1;let o=l.data;for(let i=0;i<o.length;i++){o[i].add_time=b(o[i].add_time).format("YYYY-MM-DD HH:mm:ss");for(let d=0;d<f.value.length;d++)o[i].group_id===f.value[d].id&&(o[i].group_id=f.value[d].name)}m.value===1&&(S.value=l.total),y.value=o},G=async()=>{v.value=!0;let l=(await F()).data;for(let o=0;o<l.length;o++)l[o].created_at=b(l[o].created_at).format("YYYY-MM-DD HH:mm:ss"),l[o].updated_at=b(l[o].updated_at).format("YYYY-MM-DD HH:mm:ss");f.value.push(...l)},L=()=>{m.value===1?h():m.value=1},O=s=>{if(s==null)return a.value.start_time="",a.value.end_time="",!1;if(s[0]===""||s[1]===0)return!1;a.value.start_time=s[0],a.value.end_time=s[1]};return X(m,(s,l)=>{h()}),(s,l)=>{const o=r("el-option"),i=r("el-select"),d=r("el-form-item"),C=r("el-input"),B=r("el-date-picker"),H=r("el-button"),M=r("el-form"),c=r("el-table-column"),N=r("el-table"),q=r("el-pagination"),R=r("el-card"),k=Z("loading");return p(),D(R,{shadow:"never"},{header:n(()=>[re]),default:n(()=>[w("div",se,[e(M,{model:a.value,inline:""},{default:n(()=>[e(d,null,{default:n(()=>[e(i,{filterable:"",modelValue:a.value.group_id,"onUpdate:modelValue":l[0]||(l[0]=t=>a.value.group_id=t),placeholder:"分组筛选"},{default:n(()=>[(p(!0),T(I,null,E(f.value,(t,g)=>(p(),D(o,{key:g,label:t.name,value:t.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),e(d,null,{default:n(()=>[e(i,{filterable:"",modelValue:a.value.method,"onUpdate:modelValue":l[1]||(l[1]=t=>a.value.method=t),placeholder:"Method筛选"},{default:n(()=>[(p(),T(I,null,E($,(t,g)=>e(o,{key:g,label:t.name,value:t.val},null,8,["label","value"])),64))]),_:1},8,["modelValue"])]),_:1}),e(d,null,{default:n(()=>[e(C,{modelValue:a.value.url,"onUpdate:modelValue":l[2]||(l[2]=t=>a.value.url=t),placeholder:"筛选url"},null,8,["modelValue"])]),_:1}),e(d,null,{default:n(()=>[e(C,{modelValue:a.value.ip,"onUpdate:modelValue":l[3]||(l[3]=t=>a.value.ip=t),placeholder:"筛选IP"},null,8,["modelValue"])]),_:1}),e(d,null,{default:n(()=>[e(B,{modelValue:P.value,"onUpdate:modelValue":l[4]||(l[4]=t=>P.value=t),type:"datetimerange",onChange:O,shortcuts:V(A),"value-format":"YYYY-MM-DD HH:mm","range-separator":"-","start-placeholder":"开始时间","end-placeholder":"结束时间"},null,8,["modelValue","shortcuts"])]),_:1}),e(d,null,{default:n(()=>[e(H,{type:"primary",icon:V(ee),onClick:L,plain:""},{default:n(()=>[le("搜索")]),_:1},8,["icon"])]),_:1})]),_:1},8,["model"]),e(M,{inline:""},{default:n(()=>[e(d,null,{default:n(()=>[e(H,{icon:V(ae),circle:"",onClick:h},null,8,["icon"])]),_:1})]),_:1})]),U((p(),D(N,{size:"large",border:"",data:y.value,style:{width:"100%","margin-bottom":"50px"}},{default:n(()=>[e(c,{prop:"id",label:"ID",width:"100"}),e(c,{prop:"name",label:"请求名称",width:"350"}),e(c,{prop:"url",label:"请求URL"}),e(c,{prop:"group_id",label:"所属分组",width:"200"}),e(c,{prop:"method",label:"Method",width:"150"}),e(c,{prop:"ip",label:"IP",width:"280"}),e(c,{prop:"add_time",label:"创建时间",width:"220"})]),_:1},8,["data"])),[[k,v.value]]),U((p(),T("div",ie,[e(q,{"page-size":x.value,layout:"prev, pager, next, total",total:S.value,"current-page":m.value,"onUpdate:currentPage":l[5]||(l[5]=t=>m.value=t)},null,8,["page-size","total","current-page"])])),[[k,v.value]])]),_:1})}}},fe=J(_e,[["__scopeId","data-v-c535c589"]]);export{fe as default};