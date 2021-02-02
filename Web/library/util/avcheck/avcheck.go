package avcheck

import (
	"encoding/json"
	"fmt"
	"regexp"

	"Web/library/logger"
)

func init(){
	PareJson()
}

var aVjson = make(map[string]string)

// 解析提交的内容
func PareAv(s string)(string,error){
	if len(aVjson) == 0{
		return "",fmt.Errorf("杀软检测json数据反序列化失败")
	}
	r,_ := regexp.Compile("(\\w+).exe")
	tasklist := r.FindAllString(s, -1)
	if len(tasklist) < 1{
		return "",fmt.Errorf("你提交的内容格式不正确,请检查")
	}
	result := ""
	for _,process := range tasklist{
		value, ok := aVjson[process]
		if ok {
			result += process + "    " + value + "<br/>"
		}
	}
	if result == ""{
		result = "没有检测到防护软件进程存在，可放心大胆的操作!"
	}
	return result,nil
}

// json反序列化
func PareJson(){
	err := json.Unmarshal([]byte(av_json_str),&aVjson)
	if err != nil {
		logger.WebLog.Warningf("[-] [杀软检测]  json数据反序列化失败:%s", err.Error())
	}
}

var av_json_str = `{
	"360tray.exe": "360安全卫士-实时保护",
	"360safe.exe": "360安全卫士-主程序",
	"ZhuDongFangYu.exe": "360安全卫士-主动防御",
	"360sd.exe": "360杀毒",
	"a2guard.exe": "a-squared杀毒",
	"ad-watch.exe": "Lavasoft杀毒",
	"cleaner8.exe": "The Cleaner杀毒",
	"vba32lder.exe": "vb32杀毒",
	"MongoosaGUI.exe": "Mongoosa杀毒",
	"CorantiControlCenter32.exe": "Coranti2012杀毒",
	"F-PROT.exe": "F-Prot AntiVirus",
	"CMCTrayIcon.exe": "CMC杀毒",
	"K7TSecurity.exe": "K7杀毒",
	"UnThreat.exe": "UnThreat杀毒",
	"CKSoftShiedAntivirus4.exe": "Shield Antivirus杀毒",
	"AVWatchService.exe": "VIRUSfighter杀毒",
	"ArcaTasksService.exe": "ArcaVir杀毒",
	"iptray.exe": "Immunet杀毒",
	"PSafeSysTray.exe": "PSafe杀毒",
	"nspupsvc.exe": "nProtect杀毒",
	"SpywareTerminatorShield.exe": "SpywareTerminator杀毒",
	"BKavService.exe": "Bkav杀毒",
	"MsMpEng.exe": "Microsoft Security Essentials",
	"SBAMSvc.exe": "VIPRE",
	"ccSvcHst.exe": "Norton杀毒",
	"f-secure.exe": "冰岛",
	"avp.exe": "Kaspersky",
	"KvMonXP.exe": "江民杀毒",
	"RavMonD.exe": "瑞星杀毒",
	"Mcshield.exe": "Mcafee",
	"Tbmon.exe": "Mcafee",
	"Frameworkservice.exe": "Mcafee",
	"egui.exe": "ESET NOD32",
	"ekrn.exe": "ESET NOD32",
	"eguiProxy.exe": "ESET NOD32",
	"kxetray.exe": "金山毒霸",
	"knsdtray.exe": "可牛杀毒",
	"TMBMSRV.exe": "趋势杀毒",
	"avcenter.exe": "Avira(小红伞)",
	"avguard.exe": "Avira(小红伞)",
	"avgnt.exe": "Avira(小红伞)",
	"sched.exe": "Avira(小红伞)",
	"ashDisp.exe": "Avast网络安全",
	"rtvscan.exe": "诺顿杀毒",
	"ccapp.exe": "Symantec Norton",
	"NPFMntor.exe": "Norton杀毒软件相关进程",
	"ccSetMgr.exe": "赛门铁克",
	"ccRegVfy.exe": "Norton杀毒软件自身完整性检查程序",
	"vptray.exe": "Norton病毒防火墙-盾牌图标程序",
	"ksafe.exe": "金山卫士",
	"QQPCRTP.exe": "QQ电脑管家",
	"Miner.exe": "流量矿石",
	"AYAgent.exe": "韩国胶囊",
	"patray.exe": "安博士",
	"V3Svc.exe": "安博士V3",
	"avgwdsvc.exe": "AVG杀毒",
	"QUHLPSVC.exe": "QUICK HEAL杀毒",
	"mssecess.exe": "微软杀毒",
	"SavProgress.exe": "Sophos杀毒",
	"SophosUI.exe": "Sophos杀毒",
	"SophosFS.exe": "Sophos杀毒",
	"SophosHealth.exe": "Sophos杀毒",
	"SophosSafestore64.exe": "Sophos杀毒",
	"SophosCleanM.exe": "Sophos杀毒",
	"fsavgui.exe": "F-Secure杀毒",
	"vsserv.exe": "比特梵德",
	"remupd.exe": "熊猫卫士",
	"FortiTray.exe": "飞塔",
	"safedog.exe": "安全狗",
	"parmor.exe": "木马克星",
	"Iparmor.exe.exe": "木马克星",
	"beikesan.exe": "贝壳云安全",
	"KSWebShield.exe": "金山网盾",
	"TrojanHunter.exe": "木马猎手",
	"GG.exe": "巨盾网游安全盾",
	"adam.exe": "绿鹰安全精灵",
	"AST.exe": "超级巡警",
	"ananwidget.exe": "墨者安全专家",
	"AVK.exe": "GData",
	"avg.exe": "AVG Anti-Virus",
	"spidernt.exe": "Dr.web",
	"avgaurd.exe": "Avira Antivir",
	"vsmon.exe": "ZoneAlarm",
	"cpf.exe": "Comodo",
	"outpost.exe": "Outpost Firewall",
	"rfwmain.exe": "瑞星防火墙",
	"kpfwtray.exe": "金山网镖",
	"FYFireWall.exe": "风云防火墙",
	"MPMon.exe": "微点主动防御",
	"pfw.exe": "天网防火墙",
	"S.exe": "在抓鸡",
	"1433.exe": "在扫1433",
	"DUB.exe": "在爆破",
	"ServUDaemon.exe": "发现S-U",
	"BaiduSdSvc.exe": "百度杀毒-服务进程",
	"BaiduSdTray.exe": "百度杀毒-托盘进程",
	"BaiduSd.exe": "百度杀毒-主程序",
	"SafeDogGuardCenter.exe": "安全狗",
	"safedogupdatecenter.exe": "安全狗",
	"safedogguardcenter.exe": "安全狗",
	"SafeDogSiteIIS.exe": "安全狗",
	"SafeDogTray.exe": "安全狗",
	"SafeDogServerUI.exe": "安全狗",
	"D_Safe_Manage.exe": "D盾",
	"d_manage.exe": "D盾",
	"yunsuo_agent_service.exe": "云锁",
	"yunsuo_agent_daemon.exe": "云锁",
	"HwsPanel.exe": "护卫神",
	"hws_ui.exe": "护卫神",
	"hws.exe": "护卫神",
	"hwsd.exe": "护卫神",
	"hipstray.exe": "火绒",
	"wsctrl.exe": "火绒",
	"usysdiag.exe": "火绒",
	"WEBSCANX.exe": "网络病毒克星",
	"SPHINX.exe": "SPHINX防火墙",
	"bddownloader.exe": "百度卫士",
	"baiduansvx.exe": "百度卫士-主进程",
	"AvastUI.exe": "Avast!5主程序",
	"emet_agent.exe": "已知杀软进程,名称暂未收录",
	"emet_service.exe": "已知杀软进程,名称暂未收录",
	"firesvc.exe": "已知杀软进程,名称暂未收录",
	"firetray.exe": "已知杀软进程,名称暂未收录",
	"hipsvc.exe": "已知杀软进程,名称暂未收录",
	"mfevtps.exe": "已知杀软进程,名称暂未收录",
	"mcafeefire.exe": "已知杀软进程,名称暂未收录",
	"scan32.exe": "已知杀软进程,名称暂未收录",
	"shstat.exe": "已知杀软进程,名称暂未收录",
	"vstskmgr.exe": "已知杀软进程,名称暂未收录",
	"engineserver.exe": "已知杀软进程,名称暂未收录",
	"mfeann.exe": "已知杀软进程,名称暂未收录",
	"mcscript.exe": "已知杀软进程,名称暂未收录",
	"updaterui.exe": "已知杀软进程,名称暂未收录",
	"udaterui.exe": "已知杀软进程,名称暂未收录",
	"naprdmgr.exe": "已知杀软进程,名称暂未收录",
	"cleanup.exe": "已知杀软进程,名称暂未收录",
	"cmdagent.exe": "已知杀软进程,名称暂未收录",
	"frminst.exe": "已知杀软进程,名称暂未收录",
	"mcscript_inuse.exe": "已知杀软进程,名称暂未收录",
	"mctray.exe": "已知杀软进程,名称暂未收录",
	"AAWTray.exe": "已知杀软进程,名称暂未收录",
	"Ad-Aware.exe": "已知杀软进程,名称暂未收录",
	"MSASCui.exe": "已知杀软进程,名称暂未收录",
	"_avp32.exe": "已知杀软进程,名称暂未收录",
	"_avpcc.exe": "已知杀软进程,名称暂未收录",
	"_avpm.exe": "已知杀软进程,名称暂未收录",
	"aAvgApi.exe": "已知杀软进程,名称暂未收录",
	"ackwin32.exe": "已知杀软进程,名称暂未收录",
	"adaware.exe": "已知杀软进程,名称暂未收录",
	"advxdwin.exe": "已知杀软进程,名称暂未收录",
	"agentsvr.exe": "已知杀软进程,名称暂未收录",
	"agentw.exe": "已知杀软进程,名称暂未收录",
	"alertsvc.exe": "已知杀软进程,名称暂未收录",
	"alevir.exe": "已知杀软进程,名称暂未收录",
	"alogserv.exe": "已知杀软进程,名称暂未收录",
	"amon9x.exe": "已知杀软进程,名称暂未收录",
	"anti-trojan.exe": "已知杀软进程,名称暂未收录",
	"antivirus.exe": "已知杀软进程,名称暂未收录",
	"ants.exe": "已知杀软进程,名称暂未收录",
	"apimonitor.exe": "已知杀软进程,名称暂未收录",
	"aplica32.exe": "已知杀软进程,名称暂未收录",
	"apvxdwin.exe": "已知杀软进程,名称暂未收录",
	"arr.exe": "已知杀软进程,名称暂未收录",
	"atcon.exe": "已知杀软进程,名称暂未收录",
	"atguard.exe": "已知杀软进程,名称暂未收录",
	"atro55en.exe": "已知杀软进程,名称暂未收录",
	"atupdater.exe": "已知杀软进程,名称暂未收录",
	"atwatch.exe": "已知杀软进程,名称暂未收录",
	"au.exe": "已知杀软进程,名称暂未收录",
	"aupdate.exe": "已知杀软进程,名称暂未收录",
	"auto-protect.nav80try.exe": "已知杀软进程,名称暂未收录",
	"autodown.exe": "已知杀软进程,名称暂未收录",
	"autotrace.exe": "已知杀软进程,名称暂未收录",
	"autoupdate.exe": "已知杀软进程,名称暂未收录",
	"avconsol.exe": "已知杀软进程,名称暂未收录",
	"ave32.exe": "已知杀软进程,名称暂未收录",
	"avgcc32.exe": "已知杀软进程,名称暂未收录",
	"avgctrl.exe": "已知杀软进程,名称暂未收录",
	"avgemc.exe": "已知杀软进程,名称暂未收录",
	"avgrsx.exe": "已知杀软进程,名称暂未收录",
	"avgserv.exe": "已知杀软进程,名称暂未收录",
	"avgserv9.exe": "已知杀软进程,名称暂未收录",
	"avgw.exe": "已知杀软进程,名称暂未收录",
	"avkpop.exe": "已知杀软进程,名称暂未收录",
	"avkserv.exe": "已知杀软进程,名称暂未收录",
	"avkservice.exe": "已知杀软进程,名称暂未收录",
	"avkwctl9.exe": "已知杀软进程,名称暂未收录",
	"avltmain.exe": "已知杀软进程,名称暂未收录",
	"avnt.exe": "已知杀软进程,名称暂未收录",
	"avp32.exe": "已知杀软进程,名称暂未收录",
	"avpcc.exe": "已知杀软进程,名称暂未收录",
	"avpdos32.exe": "已知杀软进程,名称暂未收录",
	"avpm.exe": "已知杀软进程,名称暂未收录",
	"avptc32.exe": "已知杀软进程,名称暂未收录",
	"avpupd.exe": "已知杀软进程,名称暂未收录",
	"avsched32.exe": "已知杀软进程,名称暂未收录",
	"avsynmgr.exe": "已知杀软进程,名称暂未收录",
	"avwin.exe": "已知杀软进程,名称暂未收录",
	"avwin95.exe": "已知杀软进程,名称暂未收录",
	"avwinnt.exe": "已知杀软进程,名称暂未收录",
	"avwupd.exe": "已知杀软进程,名称暂未收录",
	"avwupd32.exe": "已知杀软进程,名称暂未收录",
	"avwupsrv.exe": "已知杀软进程,名称暂未收录",
	"avxmonitor9x.exe": "已知杀软进程,名称暂未收录",
	"avxmonitornt.exe": "已知杀软进程,名称暂未收录",
	"avxquar.exe": "已知杀软进程,名称暂未收录",
	"backweb.exe": "已知杀软进程,名称暂未收录",
	"bargains.exe": "已知杀软进程,名称暂未收录",
	"bd_professional.exe": "已知杀软进程,名称暂未收录",
	"beagle.exe": "已知杀软进程,名称暂未收录",
	"belt.exe": "已知杀软进程,名称暂未收录",
	"bidef.exe": "已知杀软进程,名称暂未收录",
	"bidserver.exe": "已知杀软进程,名称暂未收录",
	"bipcp.exe": "已知杀软进程,名称暂未收录",
	"bipcpevalsetup.exe": "已知杀软进程,名称暂未收录",
	"bisp.exe": "已知杀软进程,名称暂未收录",
	"blackd.exe": "已知杀软进程,名称暂未收录",
	"blackice.exe": "已知杀软进程,名称暂未收录",
	"blink.exe": "已知杀软进程,名称暂未收录",
	"blss.exe": "已知杀软进程,名称暂未收录",
	"bootconf.exe": "已知杀软进程,名称暂未收录",
	"bootwarn.exe": "已知杀软进程,名称暂未收录",
	"borg2.exe": "已知杀软进程,名称暂未收录",
	"bpc.exe": "已知杀软进程,名称暂未收录",
	"brasil.exe": "已知杀软进程,名称暂未收录",
	"bs120.exe": "已知杀软进程,名称暂未收录",
	"bundle.exe": "已知杀软进程,名称暂未收录",
	"bvt.exe": "已知杀软进程,名称暂未收录",
	"ccevtmgr.exe": "已知杀软进程,名称暂未收录",
	"ccpxysvc.exe": "已知杀软进程,名称暂未收录",
	"cdp.exe": "已知杀软进程,名称暂未收录",
	"cfd.exe": "已知杀软进程,名称暂未收录",
	"cfgwiz.exe": "已知杀软进程,名称暂未收录",
	"cfiadmin.exe": "已知杀软进程,名称暂未收录",
	"cfiaudit.exe": "已知杀软进程,名称暂未收录",
	"cfinet.exe": "已知杀软进程,名称暂未收录",
	"cfinet32.exe": "已知杀软进程,名称暂未收录",
	"claw95.exe": "已知杀软进程,名称暂未收录",
	"claw95cf.exe": "已知杀软进程,名称暂未收录",
	"clean.exe": "已知杀软进程,名称暂未收录",
	"cleaner.exe": "已知杀软进程,名称暂未收录",
	"cleaner3.exe": "已知杀软进程,名称暂未收录",
	"cleanpc.exe": "已知杀软进程,名称暂未收录",
	"click.exe": "已知杀软进程,名称暂未收录",
	"cmesys.exe": "已知杀软进程,名称暂未收录",
	"cmgrdian.exe": "已知杀软进程,名称暂未收录",
	"cmon016.exe": "已知杀软进程,名称暂未收录",
	"connectionmonitor.exe": "已知杀软进程,名称暂未收录",
	"cpd.exe": "已知杀软进程,名称暂未收录",
	"cpf9x206.exe": "已知杀软进程,名称暂未收录",
	"cpfnt206.exe": "已知杀软进程,名称暂未收录",
	"ctrl.exe": "已知杀软进程,名称暂未收录",
	"cv.exe": "已知杀软进程,名称暂未收录",
	"cwnb181.exe": "已知杀软进程,名称暂未收录",
	"cwntdwmo.exe": "已知杀软进程,名称暂未收录",
	"datemanager.exe": "已知杀软进程,名称暂未收录",
	"dcomx.exe": "已知杀软进程,名称暂未收录",
	"defalert.exe": "已知杀软进程,名称暂未收录",
	"defscangui.exe": "已知杀软进程,名称暂未收录",
	"defwatch.exe": "已知杀软进程,名称暂未收录",
	"deputy.exe": "已知杀软进程,名称暂未收录",
	"divx.exe": "已知杀软进程,名称暂未收录",
	"dllcache.exe": "已知杀软进程,名称暂未收录",
	"dllreg.exe": "已知杀软进程,名称暂未收录",
	"doors.exe": "已知杀软进程,名称暂未收录",
	"dpf.exe": "已知杀软进程,名称暂未收录",
	"dpfsetup.exe": "已知杀软进程,名称暂未收录",
	"dpps2.exe": "已知杀软进程,名称暂未收录",
	"drwatson.exe": "已知杀软进程,名称暂未收录",
	"drweb32.exe": "已知杀软进程,名称暂未收录",
	"drwebupw.exe": "已知杀软进程,名称暂未收录",
	"dssagent.exe": "已知杀软进程,名称暂未收录",
	"dvp95.exe": "已知杀软进程,名称暂未收录",
	"dvp95_0.exe": "已知杀软进程,名称暂未收录",
	"ecengine.exe": "已知杀软进程,名称暂未收录",
	"efpeadm.exe": "已知杀软进程,名称暂未收录",
	"emsw.exe": "已知杀软进程,名称暂未收录",
	"ent.exe": "已知杀软进程,名称暂未收录",
	"esafe.exe": "已知杀软进程,名称暂未收录",
	"escanhnt.exe": "已知杀软进程,名称暂未收录",
	"escanv95.exe": "已知杀软进程,名称暂未收录",
	"espwatch.exe": "已知杀软进程,名称暂未收录",
	"ethereal.exe": "已知杀软进程,名称暂未收录",
	"etrustcipe.exe": "已知杀软进程,名称暂未收录",
	"evpn.exe": "已知杀软进程,名称暂未收录",
	"exantivirus-cnet.exe": "已知杀软进程,名称暂未收录",
	"exe.avxw.exe": "已知杀软进程,名称暂未收录",
	"expert.exe": "已知杀软进程,名称暂未收录",
	"explore.exe": "已知杀软进程,名称暂未收录",
	"f-agnt95.exe": "已知杀软进程,名称暂未收录",
	"f-prot95.exe": "已知杀软进程,名称暂未收录",
	"f-stopw.exe": "已知杀软进程,名称暂未收录",
	"fameh32.exe": "已知杀软进程,名称暂未收录",
	"fast.exe": "已知杀软进程,名称暂未收录",
	"fch32.exe": "已知杀软进程,名称暂未收录",
	"fih32.exe": "已知杀软进程,名称暂未收录",
	"findviru.exe": "已知杀软进程,名称暂未收录",
	"firewall.exe": "已知杀软进程,名称暂未收录",
	"fnrb32.exe": "已知杀软进程,名称暂未收录",
	"fp-win.exe": "已知杀软进程,名称暂未收录",
	"fp-win_trial.exe": "已知杀软进程,名称暂未收录",
	"fprot.exe": "已知杀软进程,名称暂未收录",
	"frw.exe": "已知杀软进程,名称暂未收录",
	"fsaa.exe": "已知杀软进程,名称暂未收录",
	"fsav.exe": "已知杀软进程,名称暂未收录",
	"fsav32.exe": "已知杀软进程,名称暂未收录",
	"fsav530stbyb.exe": "已知杀软进程,名称暂未收录",
	"fsav530wtbyb.exe": "已知杀软进程,名称暂未收录",
	"fsav95.exe": "已知杀软进程,名称暂未收录",
	"fsgk32.exe": "已知杀软进程,名称暂未收录",
	"fsm32.exe": "已知杀软进程,名称暂未收录",
	"fsma32.exe": "已知杀软进程,名称暂未收录",
	"fsmb32.exe": "已知杀软进程,名称暂未收录",
	"gator.exe": "已知杀软进程,名称暂未收录",
	"gbmenu.exe": "已知杀软进程,名称暂未收录",
	"gbpoll.exe": "已知杀软进程,名称暂未收录",
	"generics.exe": "已知杀软进程,名称暂未收录",
	"gmt.exe": "已知杀软进程,名称暂未收录",
	"guard.exe": "已知杀软进程,名称暂未收录",
	"guarddog.exe": "已知杀软进程,名称暂未收录",
	"hacktracersetup.exe": "已知杀软进程,名称暂未收录",
	"hbinst.exe": "已知杀软进程,名称暂未收录",
	"hbsrv.exe": "已知杀软进程,名称暂未收录",
	"hotactio.exe": "已知杀软进程,名称暂未收录",
	"hotpatch.exe": "已知杀软进程,名称暂未收录",
	"htlog.exe": "已知杀软进程,名称暂未收录",
	"htpatch.exe": "已知杀软进程,名称暂未收录",
	"hwpe.exe": "已知杀软进程,名称暂未收录",
	"hxdl.exe": "已知杀软进程,名称暂未收录",
	"hxiul.exe": "已知杀软进程,名称暂未收录",
	"iamapp.exe": "已知杀软进程,名称暂未收录",
	"iamserv.exe": "已知杀软进程,名称暂未收录",
	"iamstats.exe": "已知杀软进程,名称暂未收录",
	"ibmasn.exe": "已知杀软进程,名称暂未收录",
	"ibmavsp.exe": "已知杀软进程,名称暂未收录",
	"icload95.exe": "已知杀软进程,名称暂未收录",
	"icloadnt.exe": "已知杀软进程,名称暂未收录",
	"icmon.exe": "已知杀软进程,名称暂未收录",
	"icsupp95.exe": "已知杀软进程,名称暂未收录",
	"icsuppnt.exe": "已知杀软进程,名称暂未收录",
	"idle.exe": "已知杀软进程,名称暂未收录",
	"iedll.exe": "已知杀软进程,名称暂未收录",
	"iedriver.exe": "已知杀软进程,名称暂未收录",
	"iface.exe": "已知杀软进程,名称暂未收录",
	"ifw2000.exe": "已知杀软进程,名称暂未收录",
	"inetlnfo.exe": "已知杀软进程,名称暂未收录",
	"infus.exe": "已知杀软进程,名称暂未收录",
	"infwin.exe": "已知杀软进程,名称暂未收录",
	"init.exe": "已知杀软进程,名称暂未收录",
	"intdel.exe": "已知杀软进程,名称暂未收录",
	"intren.exe": "已知杀软进程,名称暂未收录",
	"iomon98.exe": "已知杀软进程,名称暂未收录",
	"istsvc.exe": "已知杀软进程,名称暂未收录",
	"jammer.exe": "已知杀软进程,名称暂未收录",
	"jdbgmrg.exe": "已知杀软进程,名称暂未收录",
	"jedi.exe": "已知杀软进程,名称暂未收录",
	"kavlite40eng.exe": "已知杀软进程,名称暂未收录",
	"kavpers40eng.exe": "已知杀软进程,名称暂未收录",
	"kavpf.exe": "已知杀软进程,名称暂未收录",
	"kazza.exe": "已知杀软进程,名称暂未收录",
	"keenvalue.exe": "已知杀软进程,名称暂未收录",
	"kerio-pf-213-en-win.exe": "已知杀软进程,名称暂未收录",
	"kerio-wrl-421-en-win.exe": "已知杀软进程,名称暂未收录",
	"kerio-wrp-421-en-win.exe": "已知杀软进程,名称暂未收录",
	"kernel32.exe": "已知杀软进程,名称暂未收录",
	"killprocesssetup161.exe": "已知杀软进程,名称暂未收录",
	"launcher.exe": "已知杀软进程,名称暂未收录",
	"ldnetmon.exe": "已知杀软进程,名称暂未收录",
	"ldpro.exe": "已知杀软进程,名称暂未收录",
	"ldpromenu.exe": "已知杀软进程,名称暂未收录",
	"ldscan.exe": "已知杀软进程,名称暂未收录",
	"lnetinfo.exe": "已知杀软进程,名称暂未收录",
	"loader.exe": "已知杀软进程,名称暂未收录",
	"localnet.exe": "已知杀软进程,名称暂未收录",
	"lockdown.exe": "已知杀软进程,名称暂未收录",
	"lockdown2000.exe": "已知杀软进程,名称暂未收录",
	"lookout.exe": "已知杀软进程,名称暂未收录",
	"lordpe.exe": "已知杀软进程,名称暂未收录",
	"lsetup.exe": "已知杀软进程,名称暂未收录",
	"luall.exe": "已知杀软进程,名称暂未收录",
	"luau.exe": "已知杀软进程,名称暂未收录",
	"lucomserver.exe": "已知杀软进程,名称暂未收录",
	"luinit.exe": "已知杀软进程,名称暂未收录",
	"luspt.exe": "已知杀软进程,名称暂未收录",
	"mapisvc32.exe": "已知杀软进程,名称暂未收录",
	"mcagent.exe": "已知杀软进程,名称暂未收录",
	"mcmnhdlr.exe": "已知杀软进程,名称暂未收录",
	"mctool.exe": "已知杀软进程,名称暂未收录",
	"mcupdate.exe": "已知杀软进程,名称暂未收录",
	"mcvsrte.exe": "已知杀软进程,名称暂未收录",
	"mcvsshld.exe": "已知杀软进程,名称暂未收录",
	"md.exe": "已知杀软进程,名称暂未收录",
	"mfin32.exe": "已知杀软进程,名称暂未收录",
	"mfw2en.exe": "已知杀软进程,名称暂未收录",
	"mfweng3.02d30.exe": "已知杀软进程,名称暂未收录",
	"mgavrtcl.exe": "已知杀软进程,名称暂未收录",
	"mgavrte.exe": "已知杀软进程,名称暂未收录",
	"mghtml.exe": "已知杀软进程,名称暂未收录",
	"mgui.exe": "已知杀软进程,名称暂未收录",
	"minilog.exe": "已知杀软进程,名称暂未收录",
	"mmod.exe": "已知杀软进程,名称暂未收录",
	"monitor.exe": "已知杀软进程,名称暂未收录",
	"moolive.exe": "已知杀软进程,名称暂未收录",
	"mostat.exe": "已知杀软进程,名称暂未收录",
	"mpfagent.exe": "已知杀软进程,名称暂未收录",
	"mpfservice.exe": "已知杀软进程,名称暂未收录",
	"mpftray.exe": "已知杀软进程,名称暂未收录",
	"mrflux.exe": "已知杀软进程,名称暂未收录",
	"msapp.exe": "已知杀软进程,名称暂未收录",
	"msbb.exe": "已知杀软进程,名称暂未收录",
	"msblast.exe": "已知杀软进程,名称暂未收录",
	"mscache.exe": "已知杀软进程,名称暂未收录",
	"msccn32.exe": "已知杀软进程,名称暂未收录",
	"mscman.exe": "已知杀软进程,名称暂未收录",
	"msconfig.exe": "已知杀软进程,名称暂未收录",
	"msdm.exe": "已知杀软进程,名称暂未收录",
	"msdos.exe": "已知杀软进程,名称暂未收录",
	"msiexec16.exe": "已知杀软进程,名称暂未收录",
	"msinfo32.exe": "已知杀软进程,名称暂未收录",
	"mslaugh.exe": "已知杀软进程,名称暂未收录",
	"msmgt.exe": "已知杀软进程,名称暂未收录",
	"msmsgri32.exe": "已知杀软进程,名称暂未收录",
	"mssmmc32.exe": "已知杀软进程,名称暂未收录",
	"mssys.exe": "已知杀软进程,名称暂未收录",
	"msvxd.exe": "已知杀软进程,名称暂未收录",
	"mu0311ad.exe": "已知杀软进程,名称暂未收录",
	"mwatch.exe": "已知杀软进程,名称暂未收录",
	"n32scanw.exe": "已知杀软进程,名称暂未收录",
	"nav.exe": "已知杀软进程,名称暂未收录",
	"navap.navapsvc.exe": "已知杀软进程,名称暂未收录",
	"navapsvc.exe": "已知杀软进程,名称暂未收录",
	"navapw32.exe": "已知杀软进程,名称暂未收录",
	"navdx.exe": "已知杀软进程,名称暂未收录",
	"navlu32.exe": "已知杀软进程,名称暂未收录",
	"navnt.exe": "已知杀软进程,名称暂未收录",
	"navstub.exe": "已知杀软进程,名称暂未收录",
	"navw32.exe": "已知杀软进程,名称暂未收录",
	"navwnt.exe": "已知杀软进程,名称暂未收录",
	"nc2000.exe": "已知杀软进程,名称暂未收录",
	"ncinst4.exe": "已知杀软进程,名称暂未收录",
	"ndd32.exe": "已知杀软进程,名称暂未收录",
	"neomonitor.exe": "已知杀软进程,名称暂未收录",
	"neowatchlog.exe": "已知杀软进程,名称暂未收录",
	"netarmor.exe": "已知杀软进程,名称暂未收录",
	"netd32.exe": "已知杀软进程,名称暂未收录",
	"netinfo.exe": "已知杀软进程,名称暂未收录",
	"netmon.exe": "已知杀软进程,名称暂未收录",
	"netscanpro.exe": "已知杀软进程,名称暂未收录",
	"netspyhunter-1.2.exe": "已知杀软进程,名称暂未收录",
	"netstat.exe": "已知杀软进程,名称暂未收录",
	"netutils.exe": "已知杀软进程,名称暂未收录",
	"nisserv.exe": "已知杀软进程,名称暂未收录",
	"nisum.exe": "已知杀软进程,名称暂未收录",
	"nmain.exe": "已知杀软进程,名称暂未收录",
	"nod32.exe": "已知杀软进程,名称暂未收录",
	"normist.exe": "已知杀软进程,名称暂未收录",
	"norton_internet_secu_3.0_407.exe": "已知杀软进程,名称暂未收录",
	"notstart.exe": "已知杀软进程,名称暂未收录",
	"npf40_tw_98_nt_me_2k.exe": "已知杀软进程,名称暂未收录",
	"npfmessenger.exe": "已知杀软进程,名称暂未收录",
	"nprotect.exe": "已知杀软进程,名称暂未收录",
	"npscheck.exe": "已知杀软进程,名称暂未收录",
	"npssvc.exe": "已知杀软进程,名称暂未收录",
	"nsched32.exe": "已知杀软进程,名称暂未收录",
	"nssys32.exe": "已知杀软进程,名称暂未收录",
	"nstask32.exe": "已知杀软进程,名称暂未收录",
	"nsupdate.exe": "已知杀软进程,名称暂未收录",
	"nt.exe": "已知杀软进程,名称暂未收录",
	"ntrtscan.exe": "已知杀软进程,名称暂未收录",
	"ntvdm.exe": "已知杀软进程,名称暂未收录",
	"ntxconfig.exe": "已知杀软进程,名称暂未收录",
	"nui.exe": "已知杀软进程,名称暂未收录",
	"nupgrade.exe": "已知杀软进程,名称暂未收录",
	"nvarch16.exe": "已知杀软进程,名称暂未收录",
	"nvc95.exe": "已知杀软进程,名称暂未收录",
	"nvsvc32.exe": "已知杀软进程,名称暂未收录",
	"nwinst4.exe": "已知杀软进程,名称暂未收录",
	"nwservice.exe": "已知杀软进程,名称暂未收录",
	"nwtool16.exe": "已知杀软进程,名称暂未收录",
	"ollydbg.exe": "已知杀软进程,名称暂未收录",
	"onsrvr.exe": "已知杀软进程,名称暂未收录",
	"optimize.exe": "已知杀软进程,名称暂未收录",
	"ostronet.exe": "已知杀软进程,名称暂未收录",
	"otfix.exe": "已知杀软进程,名称暂未收录",
	"outpostinstall.exe": "已知杀软进程,名称暂未收录",
	"outpostproinstall.exe": "已知杀软进程,名称暂未收录",
	"padmin.exe": "已知杀软进程,名称暂未收录",
	"panixk.exe": "已知杀软进程,名称暂未收录",
	"patch.exe": "已知杀软进程,名称暂未收录",
	"pavcl.exe": "已知杀软进程,名称暂未收录",
	"pavproxy.exe": "已知杀软进程,名称暂未收录",
	"pavsched.exe": "已知杀软进程,名称暂未收录",
	"pavw.exe": "已知杀软进程,名称暂未收录",
	"pccwin98.exe": "已知杀软进程,名称暂未收录",
	"pcfwallicon.exe": "已知杀软进程,名称暂未收录",
	"pcip10117_0.exe": "已知杀软进程,名称暂未收录",
	"pcscan.exe": "已知杀软进程,名称暂未收录",
	"pdsetup.exe": "已知杀软进程,名称暂未收录",
	"periscope.exe": "已知杀软进程,名称暂未收录",
	"persfw.exe": "已知杀软进程,名称暂未收录",
	"perswf.exe": "已知杀软进程,名称暂未收录",
	"pf2.exe": "已知杀软进程,名称暂未收录",
	"pfwadmin.exe": "已知杀软进程,名称暂未收录",
	"pgmonitr.exe": "已知杀软进程,名称暂未收录",
	"pingscan.exe": "已知杀软进程,名称暂未收录",
	"platin.exe": "已知杀软进程,名称暂未收录",
	"pop3trap.exe": "已知杀软进程,名称暂未收录",
	"poproxy.exe": "已知杀软进程,名称暂未收录",
	"popscan.exe": "已知杀软进程,名称暂未收录",
	"portdetective.exe": "已知杀软进程,名称暂未收录",
	"portmonitor.exe": "已知杀软进程,名称暂未收录",
	"powerscan.exe": "已知杀软进程,名称暂未收录",
	"ppinupdt.exe": "已知杀软进程,名称暂未收录",
	"pptbc.exe": "已知杀软进程,名称暂未收录",
	"ppvstop.exe": "已知杀软进程,名称暂未收录",
	"prizesurfer.exe": "已知杀软进程,名称暂未收录",
	"prmt.exe": "已知杀软进程,名称暂未收录",
	"prmvr.exe": "已知杀软进程,名称暂未收录",
	"procdump.exe": "已知杀软进程,名称暂未收录",
	"processmonitor.exe": "已知杀软进程,名称暂未收录",
	"procexplorerv1.0.exe": "已知杀软进程,名称暂未收录",
	"programauditor.exe": "已知杀软进程,名称暂未收录",
	"proport.exe": "已知杀软进程,名称暂未收录",
	"protectx.exe": "已知杀软进程,名称暂未收录",
	"pspf.exe": "已知杀软进程,名称暂未收录",
	"purge.exe": "已知杀软进程,名称暂未收录",
	"qconsole.exe": "已知杀软进程,名称暂未收录",
	"qserver.exe": "已知杀软进程,名称暂未收录",
	"rapapp.exe": "已知杀软进程,名称暂未收录",
	"rav7.exe": "已知杀软进程,名称暂未收录",
	"rav7win.exe": "已知杀软进程,名称暂未收录",
	"rav8win32eng.exe": "已知杀软进程,名称暂未收录",
	"ray.exe": "已知杀软进程,名称暂未收录",
	"rb32.exe": "已知杀软进程,名称暂未收录",
	"rcsync.exe": "已知杀软进程,名称暂未收录",
	"realmon.exe": "已知杀软进程,名称暂未收录",
	"reged.exe": "已知杀软进程,名称暂未收录",
	"regedit.exe": "已知杀软进程,名称暂未收录",
	"regedt32.exe": "已知杀软进程,名称暂未收录",
	"rescue.exe": "已知杀软进程,名称暂未收录",
	"rescue32.exe": "已知杀软进程,名称暂未收录",
	"rrguard.exe": "已知杀软进程,名称暂未收录",
	"rshell.exe": "已知杀软进程,名称暂未收录",
	"rtvscn95.exe": "已知杀软进程,名称暂未收录",
	"rulaunch.exe": "已知杀软进程,名称暂未收录",
	"run32dll.exe": "已知杀软进程,名称暂未收录",
	"rundll.exe": "已知杀软进程,名称暂未收录",
	"rundll16.exe": "已知杀软进程,名称暂未收录",
	"ruxdll32.exe": "已知杀软进程,名称暂未收录",
	"safeweb.exe": "已知杀软进程,名称暂未收录",
	"sahagent.exescan32.exe": "已知杀软进程,名称暂未收录",
	"save.exe": "已知杀软进程,名称暂未收录",
	"savenow.exe": "已知杀软进程,名称暂未收录",
	"sbserv.exe": "已知杀软进程,名称暂未收录",
	"sc.exe": "已知杀软进程,名称暂未收录",
	"scam32.exe": "已知杀软进程,名称暂未收录",
	"scan95.exe": "已知杀软进程,名称暂未收录",
	"scanpm.exe": "已知杀软进程,名称暂未收录",
	"scrscan.exe": "已知杀软进程,名称暂未收录",
	"serv95.exe": "已知杀软进程,名称暂未收录",
	"setup_flowprotector_us.exe": "已知杀软进程,名称暂未收录",
	"setupvameeval.exe": "已知杀软进程,名称暂未收录",
	"sfc.exe": "已知杀软进程,名称暂未收录",
	"sgssfw32.exe": "已知杀软进程,名称暂未收录",
	"sh.exe": "已知杀软进程,名称暂未收录",
	"shellspyinstall.exe": "已知杀软进程,名称暂未收录",
	"shn.exe": "已知杀软进程,名称暂未收录",
	"showbehind.exe": "已知杀软进程,名称暂未收录",
	"smc.exe": "已知杀软进程,名称暂未收录",
	"sms.exe": "已知杀软进程,名称暂未收录",
	"smss32.exe": "已知杀软进程,名称暂未收录",
	"soap.exe": "已知杀软进程,名称暂未收录",
	"sofi.exe": "已知杀软进程,名称暂未收录",
	"sperm.exe": "已知杀软进程,名称暂未收录",
	"spf.exe": "已知杀软进程,名称暂未收录",
	"spoler.exe": "已知杀软进程,名称暂未收录",
	"spoolcv.exe": "已知杀软进程,名称暂未收录",
	"spoolsv32.exe": "已知杀软进程,名称暂未收录",
	"spyxx.exe": "已知杀软进程,名称暂未收录",
	"srexe.exe": "已知杀软进程,名称暂未收录",
	"srng.exe": "已知杀软进程,名称暂未收录",
	"ss3edit.exe": "已知杀软进程,名称暂未收录",
	"ssg_4104.exe": "已知杀软进程,名称暂未收录",
	"ssgrate.exe": "已知杀软进程,名称暂未收录",
	"st2.exe": "已知杀软进程,名称暂未收录",
	"start.exe": "已知杀软进程,名称暂未收录",
	"stcloader.exe": "已知杀软进程,名称暂未收录",
	"supftrl.exe": "已知杀软进程,名称暂未收录",
	"support.exe": "已知杀软进程,名称暂未收录",
	"supporter5.exe": "已知杀软进程,名称暂未收录",
	"svchostc.exe": "已知杀软进程,名称暂未收录",
	"svchosts.exe": "已知杀软进程,名称暂未收录",
	"sweep95.exe": "已知杀软进程,名称暂未收录",
	"sweepnet.sweepsrv.sys.swnetsup.exe": "已知杀软进程,名称暂未收录",
	"symproxysvc.exe": "已知杀软进程,名称暂未收录",
	"symtray.exe": "已知杀软进程,名称暂未收录",
	"sysedit.exe": "已知杀软进程,名称暂未收录",
	"sysupd.exe": "已知杀软进程,名称暂未收录",
	"taskmg.exe": "已知杀软进程,名称暂未收录",
	"taskmo.exe": "已知杀软进程,名称暂未收录",
	"taumon.exe": "已知杀软进程,名称暂未收录",
	"tbscan.exe": "已知杀软进程,名称暂未收录",
	"tc.exe": "已知杀软进程,名称暂未收录",
	"tca.exe": "已知杀软进程,名称暂未收录",
	"tcm.exe": "已知杀软进程,名称暂未收录",
	"tds-3.exe": "已知杀软进程,名称暂未收录",
	"tds2-98.exe": "已知杀软进程,名称暂未收录",
	"tds2-nt.exe": "已知杀软进程,名称暂未收录",
	"teekids.exe": "已知杀软进程,名称暂未收录",
	"tfak.exe": "已知杀软进程,名称暂未收录",
	"tfak5.exe": "已知杀软进程,名称暂未收录",
	"tgbob.exe": "已知杀软进程,名称暂未收录",
	"titanin.exe": "已知杀软进程,名称暂未收录",
	"titaninxp.exe": "已知杀软进程,名称暂未收录",
	"tracert.exe": "已知杀软进程,名称暂未收录",
	"trickler.exe": "已知杀软进程,名称暂未收录",
	"trjscan.exe": "已知杀软进程,名称暂未收录",
	"trjsetup.exe": "已知杀软进程,名称暂未收录",
	"trojantrap3.exe": "已知杀软进程,名称暂未收录",
	"tsadbot.exe": "已知杀软进程,名称暂未收录",
	"tvmd.exe": "已知杀软进程,名称暂未收录",
	"tvtmd.exe": "已知杀软进程,名称暂未收录",
	"undoboot.exe": "已知杀软进程,名称暂未收录",
	"updat.exe": "已知杀软进程,名称暂未收录",
	"update.exe": "已知杀软进程,名称暂未收录",
	"upgrad.exe": "已知杀软进程,名称暂未收录",
	"utpost.exe": "已知杀软进程,名称暂未收录",
	"vbcmserv.exe": "已知杀软进程,名称暂未收录",
	"vbcons.exe": "已知杀软进程,名称暂未收录",
	"vbust.exe": "已知杀软进程,名称暂未收录",
	"vbwin9x.exe": "已知杀软进程,名称暂未收录",
	"vbwinntw.exe": "已知杀软进程,名称暂未收录",
	"vcsetup.exe": "已知杀软进程,名称暂未收录",
	"vet32.exe": "已知杀软进程,名称暂未收录",
	"vet95.exe": "已知杀软进程,名称暂未收录",
	"vettray.exe": "已知杀软进程,名称暂未收录",
	"vfsetup.exe": "已知杀软进程,名称暂未收录",
	"vir-help.exe": "已知杀软进程,名称暂未收录",
	"virusmdpersonalfirewall.exe": "已知杀软进程,名称暂未收录",
	"vnlan300.exe": "已知杀软进程,名称暂未收录",
	"vnpc3000.exe": "已知杀软进程,名称暂未收录",
	"vpc32.exe": "已知杀软进程,名称暂未收录",
	"vpc42.exe": "已知杀软进程,名称暂未收录",
	"vpfw30s.exe": "已知杀软进程,名称暂未收录",
	"vscan40.exe": "已知杀软进程,名称暂未收录",
	"vscenu6.02d30.exe": "已知杀软进程,名称暂未收录",
	"vsched.exe": "已知杀软进程,名称暂未收录",
	"vsecomr.exe": "已知杀软进程,名称暂未收录",
	"vshwin32.exe": "已知杀软进程,名称暂未收录",
	"vsisetup.exe": "已知杀软进程,名称暂未收录",
	"vsmain.exe": "已知杀软进程,名称暂未收录",
	"vsstat.exe": "已知杀软进程,名称暂未收录",
	"vswin9xe.exe": "已知杀软进程,名称暂未收录",
	"vswinntse.exe": "已知杀软进程,名称暂未收录",
	"vswinperse.exe": "已知杀软进程,名称暂未收录",
	"w32dsm89.exe": "已知杀软进程,名称暂未收录",
	"w9x.exe": "已知杀软进程,名称暂未收录",
	"watchdog.exe": "已知杀软进程,名称暂未收录",
	"webdav.exe": "已知杀软进程,名称暂未收录",
	"webtrap.exe": "已知杀软进程,名称暂未收录",
	"wfindv32.exe": "已知杀软进程,名称暂未收录",
	"whoswatchingme.exe": "已知杀软进程,名称暂未收录",
	"wimmun32.exe": "已知杀软进程,名称暂未收录",
	"win-bugsfix.exe": "已知杀软进程,名称暂未收录",
	"win32.exe": "已知杀软进程,名称暂未收录",
	"win32us.exe": "已知杀软进程,名称暂未收录",
	"winactive.exe": "已知杀软进程,名称暂未收录",
	"window.exe": "已知杀软进程,名称暂未收录",
	"windows.exe": "已知杀软进程,名称暂未收录",
	"wininetd.exe": "已知杀软进程,名称暂未收录",
	"wininitx.exe": "已知杀软进程,名称暂未收录",
	"winlogin.exe": "已知杀软进程,名称暂未收录",
	"winmain.exe": "已知杀软进程,名称暂未收录",
	"winnet.exe": "已知杀软进程,名称暂未收录",
	"winppr32.exe": "已知杀软进程,名称暂未收录",
	"winrecon.exe": "已知杀软进程,名称暂未收录",
	"winservn.exe": "已知杀软进程,名称暂未收录",
	"winssk32.exe": "已知杀软进程,名称暂未收录",
	"winstart.exe": "已知杀软进程,名称暂未收录",
	"winstart001.exe": "已知杀软进程,名称暂未收录",
	"wintsk32.exe": "已知杀软进程,名称暂未收录",
	"winupdate.exe": "已知杀软进程,名称暂未收录",
	"wkufind.exe": "已知杀软进程,名称暂未收录",
	"wnad.exe": "已知杀软进程,名称暂未收录",
	"wnt.exe": "已知杀软进程,名称暂未收录",
	"wradmin.exe": "已知杀软进程,名称暂未收录",
	"wrctrl.exe": "已知杀软进程,名称暂未收录",
	"wsbgate.exe": "已知杀软进程,名称暂未收录",
	"wupdater.exe": "已知杀软进程,名称暂未收录",
	"wupdt.exe": "已知杀软进程,名称暂未收录",
	"wyvernworksfirewall.exe": "已知杀软进程,名称暂未收录",
	"xpf202en.exe": "已知杀软进程,名称暂未收录",
	"zapro.exe": "已知杀软进程,名称暂未收录",
	"zapsetup3001.exe": "已知杀软进程,名称暂未收录",
	"zatutor.exe": "已知杀软进程,名称暂未收录",
	"zonalm2601.exe": "已知杀软进程,名称暂未收录",
	"zonealarm.exe": "已知杀软进程,名称暂未收录",
	"AVPM.exe": "已知杀软进程,名称暂未收录",
	"A2CMD.exe": "已知杀软进程,名称暂未收录",
	"A2SERVICE.exe": "已知杀软进程,名称暂未收录",
	"A2FREE.exe": "已知杀软进程,名称暂未收录",
	"ADVCHK.exe": "已知杀软进程,名称暂未收录",
	"AGB.exe": "已知杀软进程,名称暂未收录",
	"AKRNL.exe": "已知杀软进程,名称暂未收录",
	"AHPROCMONSERVER.exe": "已知杀软进程,名称暂未收录",
	"AIRDEFENSE.exe": "已知杀软进程,名称暂未收录",
	"ALERTSVC.exe": "已知杀软进程,名称暂未收录",
	"AVIRA.exe": "已知杀软进程,名称暂未收录",
	"AMON.exe": "已知杀软进程,名称暂未收录",
	"TROJAN.exe": "已知杀软进程,名称暂未收录",
	"AVZ.exe": "已知杀软进程,名称暂未收录",
	"ANTIVIR.exe": "已知杀软进程,名称暂未收录",
	"APVXDWIN.exe": "已知杀软进程,名称暂未收录",
	"ARMOR2NET.exe": "已知杀软进程,名称暂未收录",
	"ASH.exeexe.exe": "已知杀软进程,名称暂未收录",
	"ASHENHCD.exe": "已知杀软进程,名称暂未收录",
	"ASHMAISV.exe": "已知杀软进程,名称暂未收录",
	"ASHPOPWZ.exe": "已知杀软进程,名称暂未收录",
	"ASHSERV.exe": "已知杀软进程,名称暂未收录",
	"ASHSIMPL.exe": "已知杀软进程,名称暂未收录",
	"ASHSKPCK.exe": "已知杀软进程,名称暂未收录",
	"ASHWEBSV.exe": "已知杀软进程,名称暂未收录",
	"ASWUPDSV.exe": "已知杀软进程,名称暂未收录",
	"ASWSCAN.exe": "已知杀软进程,名称暂未收录",
	"AVCIMAN.exe": "已知杀软进程,名称暂未收录",
	"AVCONSOL.exe": "已知杀软进程,名称暂未收录",
	"AVENGINE.exe": "已知杀软进程,名称暂未收录",
	"AVESVC.exe": "已知杀软进程,名称暂未收录",
	"AVEVAL.exe": "已知杀软进程,名称暂未收录",
	"AVEVL32.exe": "已知杀软进程,名称暂未收录",
	"AVGAM.exe": "已知杀软进程,名称暂未收录",
	"AVGCC.exe": "已知杀软进程,名称暂未收录",
	"AVGCHSVX.exe": "已知杀软进程,名称暂未收录",
	"AVGCSRVX.exe": "已知杀软进程,名称暂未收录",
	"AVGNSX.exe": "已知杀软进程,名称暂未收录",
	"AVGCC32.exe": "已知杀软进程,名称暂未收录",
	"AVGCTRL.exe": "已知杀软进程,名称暂未收录",
	"AVGEMC.exe": "已知杀软进程,名称暂未收录",
	"AVGFWSRV.exe": "已知杀软进程,名称暂未收录",
	"AVGNTMGR.exe": "已知杀软进程,名称暂未收录",
	"AVGSERV.exe": "已知杀软进程,名称暂未收录",
	"AVGTRAY.exe": "已知杀软进程,名称暂未收录",
	"AVGUPSVC.exe": "已知杀软进程,名称暂未收录",
	"AVINITNT.exe": "已知杀软进程,名称暂未收录",
	"AVKSERV.exe": "已知杀软进程,名称暂未收录",
	"AVKSERVICE.exe": "已知杀软进程,名称暂未收录",
	"AVKWCTL.exe": "已知杀软进程,名称暂未收录",
	"AVP32.exe": "已知杀软进程,名称暂未收录",
	"AVPCC.exe": "已知杀软进程,名称暂未收录",
	"AVSERVER.exe": "已知杀软进程,名称暂未收录",
	"AVSCHED32.exe": "已知杀软进程,名称暂未收录",
	"AVSYNMGR.exe": "已知杀软进程,名称暂未收录",
	"AVWUPD32.exe": "已知杀软进程,名称暂未收录",
	"AVWUPSRV.exe": "已知杀软进程,名称暂未收录",
	"AVXMONITOR.exe": "已知杀软进程,名称暂未收录",
	"AVXQUAR.exe": "已知杀软进程,名称暂未收录",
	"BDSWITCH.exe": "已知杀软进程,名称暂未收录",
	"BLACKD.exe": "已知杀软进程,名称暂未收录",
	"BLACKICE.exe": "已知杀软进程,名称暂未收录",
	"CAFIX.exe": "已知杀软进程,名称暂未收录",
	"BITDEFENDER.exe": "已知杀软进程,名称暂未收录",
	"CCEVTMGR.exe": "已知杀软进程,名称暂未收录",
	"CFP.exe": "已知杀软进程,名称暂未收录",
	"CFPCONFIG.exe": "已知杀软进程,名称暂未收录",
	"CFIAUDIT.exe": "已知杀软进程,名称暂未收录",
	"CLAMTRAY.exe": "已知杀软进程,名称暂未收录",
	"CLAMWIN.exe": "已知杀软进程,名称暂未收录",
	"CUREIT.exe": "已知杀软进程,名称暂未收录",
	"DEFWATCH.exe": "已知杀软进程,名称暂未收录",
	"DRVIRUS.exe": "已知杀软进程,名称暂未收录",
	"DRWADINS.exe": "已知杀软进程,名称暂未收录",
	"DRWEB.exe": "已知杀软进程,名称暂未收录",
	"DEFENDERDAEMON.exe": "已知杀软进程,名称暂未收录",
	"DWEBLLIO.exe": "已知杀软进程,名称暂未收录",
	"DWEBIO.exe": "已知杀软进程,名称暂未收录",
	"ESCANH95.exe": "已知杀软进程,名称暂未收录",
	"ESCANHNT.exe": "已知杀软进程,名称暂未收录",
	"EWIDOCTRL.exe": "已知杀软进程,名称暂未收录",
	"EZANTIVIRUSREGISTRATIONCHECK.exe": "已知杀软进程,名称暂未收录",
	"F-AGNT95.exe": "已知杀软进程,名称暂未收录",
	"FAMEH32.exe": "已知杀软进程,名称暂未收录",
	"FILEMON.exe": "已知杀软进程,名称暂未收录",
	"FIREWALL.exe": "已知杀软进程,名称暂未收录",
	"FORTICLIENT.exe": "已知杀软进程,名称暂未收录",
	"FORTISCAN.exe": "已知杀软进程,名称暂未收录",
	"FPAVSERVER.exe": "已知杀软进程,名称暂未收录",
	"FPROTTRAY.exe": "已知杀软进程,名称暂未收录",
	"FPWIN.exe": "已知杀软进程,名称暂未收录",
	"FRESHCLAM.exe": "已知杀软进程,名称暂未收录",
	"FSAV32.exe": "已知杀软进程,名称暂未收录",
	"FSBWSYS.exe": "已知杀软进程,名称暂未收录",
	"F-SCHED.exe": "已知杀软进程,名称暂未收录",
	"FSDFWD.exe": "已知杀软进程,名称暂未收录",
	"FSGK32.exe": "已知杀软进程,名称暂未收录",
	"FSGK32ST.exe": "已知杀软进程,名称暂未收录",
	"FSGUIEXE.exe": "已知杀软进程,名称暂未收录",
	"FSMA32.exe": "已知杀软进程,名称暂未收录",
	"FSMB32.exe": "已知杀软进程,名称暂未收录",
	"FSPEX.exe": "已知杀软进程,名称暂未收录",
	"FSSM32.exe": "已知杀软进程,名称暂未收录",
	"F-STOPW.exe": "已知杀软进程,名称暂未收录",
	"GCASDTSERV.exe": "已知杀软进程,名称暂未收录",
	"GCASSERV.exe": "已知杀软进程,名称暂未收录",
	"GIANTANTISPYWARE.exe": "已知杀软进程,名称暂未收录",
	"GUARDGUI.exe": "已知杀软进程,名称暂未收录",
	"GUARDNT.exe": "已知杀软进程,名称暂未收录",
	"GUARDXSERVICE.exe": "已知杀软进程,名称暂未收录",
	"GUARDXKICKOFF.exe": "已知杀软进程,名称暂未收录",
	"HREGMON.exe": "已知杀软进程,名称暂未收录",
	"HRRES.exe": "已知杀软进程,名称暂未收录",
	"HSOCKPE.exe": "已知杀软进程,名称暂未收录",
	"HUPDATE.exe": "已知杀软进程,名称暂未收录",
	"IAMAPP.exe": "已知杀软进程,名称暂未收录",
	"IAMSERV.exe": "已知杀软进程,名称暂未收录",
	"ICLOAD95.exe": "已知杀软进程,名称暂未收录",
	"ICLOADNT.exe": "已知杀软进程,名称暂未收录",
	"ICMON.exe": "已知杀软进程,名称暂未收录",
	"ICSSUPPNT.exe": "已知杀软进程,名称暂未收录",
	"ICSUPP95.exe": "已知杀软进程,名称暂未收录",
	"ICSUPPNT.exe": "已知杀软进程,名称暂未收录",
	"INETUPD.exe": "已知杀软进程,名称暂未收录",
	"INOCIT.exe": "已知杀软进程,名称暂未收录",
	"INORPC.exe": "已知杀软进程,名称暂未收录",
	"INORT.exe": "已知杀软进程,名称暂未收录",
	"INOTASK.exe": "已知杀软进程,名称暂未收录",
	"INOUPTNG.exe": "已知杀软进程,名称暂未收录",
	"IOMON98.exe": "已知杀软进程,名称暂未收录",
	"ISAFE.exe": "已知杀软进程,名称暂未收录",
	"ISATRAY.exe": "已知杀软进程,名称暂未收录",
	"KAV.exe": "已知杀软进程,名称暂未收录",
	"KAVMM.exe": "已知杀软进程,名称暂未收录",
	"KAVPF.exe": "已知杀软进程,名称暂未收录",
	"KAVPFW.exe": "已知杀软进程,名称暂未收录",
	"KAVSTART.exe": "已知杀软进程,名称暂未收录",
	"KAVSVC.exe": "已知杀软进程,名称暂未收录",
	"KAVSVCUI.exe": "已知杀软进程,名称暂未收录",
	"KMAILMON.exe": "已知杀软进程,名称暂未收录",
	"MAMUTU.exe": "已知杀软进程,名称暂未收录",
	"MCAGENT.exe": "已知杀软进程,名称暂未收录",
	"MCMNHDLR.exe": "已知杀软进程,名称暂未收录",
	"MCREGWIZ.exe": "已知杀软进程,名称暂未收录",
	"MCUPDATE.exe": "已知杀软进程,名称暂未收录",
	"MCVSSHLD.exe": "已知杀软进程,名称暂未收录",
	"MINILOG.exe": "已知杀软进程,名称暂未收录",
	"MYAGTSVC.exe": "已知杀软进程,名称暂未收录",
	"MYAGTTRY.exe": "已知杀软进程,名称暂未收录",
	"NAVAPSVC.exe": "已知杀软进程,名称暂未收录",
	"NAVAPW32.exe": "已知杀软进程,名称暂未收录",
	"NAVLU32.exe": "已知杀软进程,名称暂未收录",
	"NAVW32.exe": "已知杀软进程,名称暂未收录",
	"NEOWATCHLOG.exe": "已知杀软进程,名称暂未收录",
	"NEOWATCHTRAY.exe": "已知杀软进程,名称暂未收录",
	"NISSERV.exe": "已知杀软进程,名称暂未收录",
	"NISUM.exe": "已知杀软进程,名称暂未收录",
	"NMAIN.exe": "已知杀软进程,名称暂未收录",
	"NOD32.exe": "已知杀软进程,名称暂未收录",
	"NORMIST.exe": "已知杀软进程,名称暂未收录",
	"NOTSTART.exe": "已知杀软进程,名称暂未收录",
	"NPAVTRAY.exe": "已知杀软进程,名称暂未收录",
	"NPFMSG.exe": "已知杀软进程,名称暂未收录",
	"NPROTECT.exe": "已知杀软进程,名称暂未收录",
	"NSCHED32.exe": "已知杀软进程,名称暂未收录",
	"NSMDTR.exe": "已知杀软进程,名称暂未收录",
	"NSSSERV.exe": "已知杀软进程,名称暂未收录",
	"NSSTRAY.exe": "已知杀软进程,名称暂未收录",
	"NTRTSCAN.exe": "已知杀软进程,名称暂未收录",
	"NTOS.exe": "已知杀软进程,名称暂未收录",
	"NTXCONFIG.exe": "已知杀软进程,名称暂未收录",
	"NUPGRADE.exe": "已知杀软进程,名称暂未收录",
	"NVCOD.exe": "已知杀软进程,名称暂未收录",
	"NVCTE.exe": "已知杀软进程,名称暂未收录",
	"NVCUT.exe": "已知杀软进程,名称暂未收录",
	"NWSERVICE.exe": "已知杀软进程,名称暂未收录",
	"OFCPFWSVC.exe": "已知杀软进程,名称暂未收录",
	"ONLINENT.exe": "已知杀软进程,名称暂未收录",
	"OPSSVC.exe": "已知杀软进程,名称暂未收录",
	"OP_MON.exe": "已知杀软进程,名称暂未收录",
	"PAVFIRES.exe": "已知杀软进程,名称暂未收录",
	"PAVFNSVR.exe": "已知杀软进程,名称暂未收录",
	"PAVKRE.exe": "已知杀软进程,名称暂未收录",
	"PAVPROT.exe": "已知杀软进程,名称暂未收录",
	"PAVPROXY.exe": "已知杀软进程,名称暂未收录",
	"PAVPRSRV.exe": "已知杀软进程,名称暂未收录",
	"PAVSRV51.exe": "已知杀软进程,名称暂未收录",
	"PAVSS.exe": "已知杀软进程,名称暂未收录",
	"PCCGUIDE.exe": "已知杀软进程,名称暂未收录",
	"PCCIOMON.exe": "已知杀软进程,名称暂未收录",
	"PCCNTMON.exe": "已知杀软进程,名称暂未收录",
	"PCCPFW.exe": "已知杀软进程,名称暂未收录",
	"PCCTLCOM.exe": "已知杀软进程,名称暂未收录",
	"PCTAV.exe": "已知杀软进程,名称暂未收录",
	"PERSFW.exe": "已知杀软进程,名称暂未收录",
	"PERTSK.exe": "已知杀软进程,名称暂未收录",
	"PERVAC.exe": "已知杀软进程,名称暂未收录",
	"PESTPATROL.exe": "已知杀软进程,名称暂未收录",
	"PNMSRV.exe": "已知杀软进程,名称暂未收录",
	"PREVSRV.exe": "已知杀软进程,名称暂未收录",
	"PREVX.exe": "已知杀软进程,名称暂未收录",
	"PSIMSVC.exe": "已知杀软进程,名称暂未收录",
	"QHONLINE.exe": "已知杀软进程,名称暂未收录",
	"QHONSVC.exe": "已知杀软进程,名称暂未收录",
	"QHWSCSVC.exe": "已知杀软进程,名称暂未收录",
	"QHSET.exe": "已知杀软进程,名称暂未收录",
	"RTVSCN95.exe": "已知杀软进程,名称暂未收录",
	"SALITY.exe": "已知杀软进程,名称暂未收录",
	"SAPISSVC.exe": "已知杀软进程,名称暂未收录",
	"SCANWSCS.exe": "已知杀软进程,名称暂未收录",
	"SAVADMINSERVICE.exe": "已知杀软进程,名称暂未收录",
	"SAVMAIN.exe": "已知杀软进程,名称暂未收录",
	"SAVSCAN.exe": "已知杀软进程,名称暂未收录",
	"SCANNINGPROCESS.exe": "已知杀软进程,名称暂未收录",
	"SDRA64.exe": "已知杀软进程,名称暂未收录",
	"SDHELP.exe": "已知杀软进程,名称暂未收录",
	"SHSTAT.exe": "已知杀软进程,名称暂未收录",
	"SITECLI.exe": "已知杀软进程,名称暂未收录",
	"SPBBCSVC.exe": "已知杀软进程,名称暂未收录",
	"SPIDERCPL.exe": "已知杀软进程,名称暂未收录",
	"SPIDERML.exe": "已知杀软进程,名称暂未收录",
	"SPIDERUI.exe": "已知杀软进程,名称暂未收录",
	"SPYBOTSD.exe": "已知杀软进程,名称暂未收录",
	"SPYXX.exe": "已知杀软进程,名称暂未收录",
	"SS3EDIT.exe": "已知杀软进程,名称暂未收录",
	"STOPSIGNAV.exe": "已知杀软进程,名称暂未收录",
	"SWAGENT.exe": "已知杀软进程,名称暂未收录",
	"SWDOCTOR.exe": "已知杀软进程,名称暂未收录",
	"SWNETSUP.exe": "已知杀软进程,名称暂未收录",
	"SYMLCSVC.exe": "已知杀软进程,名称暂未收录",
	"SYMPROXYSVC.exe": "已知杀软进程,名称暂未收录",
	"SYMSPORT.exe": "已知杀软进程,名称暂未收录",
	"SYMWSC.exe": "已知杀软进程,名称暂未收录",
	"SYNMGR.exe": "已知杀软进程,名称暂未收录",
	"TAUMON.exe": "已知杀软进程,名称暂未收录",
	"TMLISTEN.exe": "已知杀软进程,名称暂未收录",
	"TMNTSRV.exe": "已知杀软进程,名称暂未收录",
	"TMPROXY.exe": "已知杀软进程,名称暂未收录",
	"TNBUTIL.exe": "已知杀软进程,名称暂未收录",
	"TRJSCAN.exe": "已知杀软进程,名称暂未收录",
	"VBA32ECM.exe": "已知杀软进程,名称暂未收录",
	"VBA32IFS.exe": "已知杀软进程,名称暂未收录",
	"VBA32LDR.exe": "已知杀软进程,名称暂未收录",
	"VBA32PP3.exe": "已知杀软进程,名称暂未收录",
	"VBSNTW.exe": "已知杀软进程,名称暂未收录",
	"VCRMON.exe": "已知杀软进程,名称暂未收录",
	"VRFWSVC.exe": "已知杀软进程,名称暂未收录",
	"VRMONNT.exe": "已知杀软进程,名称暂未收录",
	"VRMONSVC.exe": "已知杀软进程,名称暂未收录",
	"VRRW32.exe": "已知杀软进程,名称暂未收录",
	"VSECOMR.exe": "已知杀软进程,名称暂未收录",
	"VSHWIN32.exe": "已知杀软进程,名称暂未收录",
	"VSSTAT.exe": "已知杀软进程,名称暂未收录",
	"WATCHDOG.exe": "已知杀软进程,名称暂未收录",
	"WINSSNOTIFY.exe": "已知杀软进程,名称暂未收录",
	"WRCTRL.exe": "已知杀软进程,名称暂未收录",
	"XCOMMSVR.exe": "已知杀软进程,名称暂未收录",
	"ZLCLIENT.exe": "已知杀软进程,名称暂未收录",
	"ZONEALARM.exe": "已知杀软进程,名称暂未收录",
	"360rp.exe": "已知杀软进程,名称暂未收录",
	"afwServ.exe": "已知杀软进程,名称暂未收录",
	"safeboxTray.exe": "已知杀软进程,名称暂未收录",
	"360safebox.exe": "已知杀软进程,名称暂未收录",
	"QQPCTray.exe": "已知杀软进程,名称暂未收录",
	"KSafeTray.exe": "已知杀软进程,名称暂未收录",
	"KSafeSvc.exe": "已知杀软进程,名称暂未收录",
	"KWatch.exe": "已知杀软进程,名称暂未收录"
}`