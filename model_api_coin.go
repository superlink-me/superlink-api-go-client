/*
Superlink

API for Superlink

API version: v0.1.17
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
	"fmt"
)

// ApiCoin the model 'ApiCoin'
type ApiCoin string

// List of api.Coin
const (
	BTC ApiCoin = "BTC"
	LTC ApiCoin = "LTC"
	DOGE ApiCoin = "DOGE"
	RDD ApiCoin = "RDD"
	DASH ApiCoin = "DASH"
	PPC ApiCoin = "PPC"
	NMC ApiCoin = "NMC"
	FTC ApiCoin = "FTC"
	XCP ApiCoin = "XCP"
	BLK ApiCoin = "BLK"
	NSR ApiCoin = "NSR"
	NBT ApiCoin = "NBT"
	MZC ApiCoin = "MZC"
	VIA ApiCoin = "VIA"
	RBY ApiCoin = "RBY"
	GRS ApiCoin = "GRS"
	DGC ApiCoin = "DGC"
	DGB ApiCoin = "DGB"
	MONA ApiCoin = "MONA"
	CLAM ApiCoin = "CLAM"
	XPM ApiCoin = "XPM"
	NEOS ApiCoin = "NEOS"
	JBS ApiCoin = "JBS"
	ZRC ApiCoin = "ZRC"
	VTC ApiCoin = "VTC"
	NXT ApiCoin = "NXT"
	BURST ApiCoin = "BURST"
	MUE ApiCoin = "MUE"
	ZOOM ApiCoin = "ZOOM"
	SDC ApiCoin = "SDC"
	PKB ApiCoin = "PKB"
	PND ApiCoin = "PND"
	START ApiCoin = "START"
	MOIN ApiCoin = "MOIN"
	EXP ApiCoin = "EXP"
	DCR ApiCoin = "DCR"
	XEM ApiCoin = "XEM"
	PART ApiCoin = "PART"
	SHR ApiCoin = "SHR"
	NVC ApiCoin = "NVC"
	AC ApiCoin = "AC"
	BTCD ApiCoin = "BTCD"
	DOPE ApiCoin = "DOPE"
	TPC ApiCoin = "TPC"
	AIB ApiCoin = "AIB"
	EDRC ApiCoin = "EDRC"
	SYS ApiCoin = "SYS"
	SLR ApiCoin = "SLR"
	SMLY ApiCoin = "SMLY"
	ETH ApiCoin = "ETH"
	PSB ApiCoin = "PSB"
	XBC ApiCoin = "XBC"
	NXS ApiCoin = "NXS"
	INSN ApiCoin = "INSN"
	OK ApiCoin = "OK"
	BRIT ApiCoin = "BRIT"
	CMP ApiCoin = "CMP"
	CRW ApiCoin = "CRW"
	BELA ApiCoin = "BELA"
	ICX ApiCoin = "ICX"
	FJC ApiCoin = "FJC"
	MIX ApiCoin = "MIX"
	CLUB ApiCoin = "CLUB"
	RICHX ApiCoin = "RICHX"
	POT ApiCoin = "POT"
	QRK ApiCoin = "QRK"
	TRC ApiCoin = "TRC"
	GRC ApiCoin = "GRC"
	AUR ApiCoin = "AUR"
	IXC ApiCoin = "IXC"
	NLG ApiCoin = "NLG"
	BITB ApiCoin = "BITB"
	XMY ApiCoin = "XMY"
	BSD ApiCoin = "BSD"
	UNO ApiCoin = "UNO"
	GB ApiCoin = "GB"
	SHM ApiCoin = "SHM"
	CRX ApiCoin = "CRX"
	BIQ ApiCoin = "BIQ"
	EVO ApiCoin = "EVO"
	STO ApiCoin = "STO"
	BIGUP ApiCoin = "BIGUP"
	GAME ApiCoin = "GAME"
	DLC ApiCoin = "DLC"
	ZYD ApiCoin = "ZYD"
	DBIC ApiCoin = "DBIC"
	STRAT ApiCoin = "STRAT"
	SH ApiCoin = "SH"
	MARS ApiCoin = "MARS"
	UBQ ApiCoin = "UBQ"
	PTC ApiCoin = "PTC"
	NRO ApiCoin = "NRO"
	ARK ApiCoin = "ARK"
	USC ApiCoin = "USC"
	THC ApiCoin = "THC"
	LINX ApiCoin = "LINX"
	ECN ApiCoin = "ECN"
	DNR ApiCoin = "DNR"
	PINK ApiCoin = "PINK"
	ATOM ApiCoin = "ATOM"
	PIVX ApiCoin = "PIVX"
	FLASH ApiCoin = "FLASH"
	ZEN ApiCoin = "ZEN"
	PUT ApiCoin = "PUT"
	ZNY ApiCoin = "ZNY"
	UNIFY ApiCoin = "UNIFY"
	XST ApiCoin = "XST"
	VC ApiCoin = "VC"
	XMR ApiCoin = "XMR"
	VOX ApiCoin = "VOX"
	NAV ApiCoin = "NAV"
	ZEC ApiCoin = "ZEC"
	LSK ApiCoin = "LSK"
	STEEM ApiCoin = "STEEM"
	XZC ApiCoin = "XZC"
	RBTC ApiCoin = "RBTC"
	RPT ApiCoin = "RPT"
	KMD ApiCoin = "KMD"
	RIC ApiCoin = "RIC"
	XRP ApiCoin = "XRP"
	NEBL ApiCoin = "NEBL"
	ZCL ApiCoin = "ZCL"
	WHL ApiCoin = "WHL"
	ERC ApiCoin = "ERC"
	DMD ApiCoin = "DMD"
	BTM ApiCoin = "BTM"
	BIO ApiCoin = "BIO"
	SSN ApiCoin = "SSN"
	TOA ApiCoin = "TOA"
	BTX ApiCoin = "BTX"
	ACC ApiCoin = "ACC"
	ELLA ApiCoin = "ELLA"
	PIRL ApiCoin = "PIRL"
	XNO ApiCoin = "XNO"
	VIVO ApiCoin = "VIVO"
	FRST ApiCoin = "FRST"
	HNC ApiCoin = "HNC"
	BUZZ ApiCoin = "BUZZ"
	MBRS ApiCoin = "MBRS"
	HC ApiCoin = "HC"
	HTML ApiCoin = "HTML"
	ODN ApiCoin = "ODN"
	ONX ApiCoin = "ONX"
	RVN ApiCoin = "RVN"
	GBX ApiCoin = "GBX"
	BTCZ ApiCoin = "BTCZ"
	POA ApiCoin = "POA"
	NYC ApiCoin = "NYC"
	MXT ApiCoin = "MXT"
	WC ApiCoin = "WC"
	MNX ApiCoin = "MNX"
	MUSIC ApiCoin = "MUSIC"
	CRAVE ApiCoin = "CRAVE"
	STAK ApiCoin = "STAK"
	LCH ApiCoin = "LCH"
	EXCL ApiCoin = "EXCL"
	LCC ApiCoin = "LCC"
	XFE ApiCoin = "XFE"
	EOS ApiCoin = "EOS"
	TRX ApiCoin = "TRX"
	KOBO ApiCoin = "KOBO"
	HUSH ApiCoin = "HUSH"
	BAN ApiCoin = "BAN"
	ETF ApiCoin = "ETF"
	OMNI ApiCoin = "OMNI"
	BIFI ApiCoin = "BIFI"
	CNMC ApiCoin = "CNMC"
	BCN ApiCoin = "BCN"
	RIN ApiCoin = "RIN"
	ATP ApiCoin = "ATP"
	EVT ApiCoin = "EVT"
	ATN ApiCoin = "ATN"
	BIS ApiCoin = "BIS"
	NEET ApiCoin = "NEET"
	BOPO ApiCoin = "BOPO"
	OOT ApiCoin = "OOT"
	ALIAS ApiCoin = "ALIAS"
	BOXY ApiCoin = "BOXY"
	FLO ApiCoin = "FLO"
	MEC ApiCoin = "MEC"
	BTDX ApiCoin = "BTDX"
	XAX ApiCoin = "XAX"
	ANON ApiCoin = "ANON"
	LTZ ApiCoin = "LTZ"
	SMART ApiCoin = "SMART"
	XUEZ ApiCoin = "XUEZ"
	HLM ApiCoin = "HLM"
	WEB ApiCoin = "WEB"
	ACM ApiCoin = "ACM"
	BITC ApiCoin = "BITC"
	TZC ApiCoin = "TZC"
	VAR ApiCoin = "VAR"
	IOV ApiCoin = "IOV"
	FIO ApiCoin = "FIO"
	BSV ApiCoin = "BSV"
	DXN ApiCoin = "DXN"
	PCX ApiCoin = "PCX"
	LOKI ApiCoin = "LOKI"
	NIM ApiCoin = "NIM"
	EXOS ApiCoin = "EXOS"
	ECA ApiCoin = "ECA"
	SOOM ApiCoin = "SOOM"
	FREE ApiCoin = "FREE"
	NPW ApiCoin = "NPW"
	BST ApiCoin = "BST"
	ZEST ApiCoin = "ZEST"
	ABT ApiCoin = "ABT"
	PION ApiCoin = "PION"
	ZBUX ApiCoin = "ZBUX"
	KPL ApiCoin = "KPL"
	TPAY ApiCoin = "TPAY"
	ZILLA ApiCoin = "ZILLA"
	ANK ApiCoin = "ANK"
	BCC ApiCoin = "BCC"
	HPB ApiCoin = "HPB"
	ONE ApiCoin = "ONE"
	SBC ApiCoin = "SBC"
	IPC ApiCoin = "IPC"
	DMTC ApiCoin = "DMTC"
	OGC ApiCoin = "OGC"
	SHIT ApiCoin = "SHIT"
	ANDES ApiCoin = "ANDES"
	AREPA ApiCoin = "AREPA"
	BOLI ApiCoin = "BOLI"
	RIL ApiCoin = "RIL"
	BRAVO ApiCoin = "BRAVO"
	ALGO ApiCoin = "ALGO"
	BZX ApiCoin = "BZX"
	GXX ApiCoin = "GXX"
	HEAT ApiCoin = "HEAT"
	XDN ApiCoin = "XDN"
	FSN ApiCoin = "FSN"
	BOLD ApiCoin = "BOLD"
	IOST ApiCoin = "IOST"
	TKEY ApiCoin = "TKEY"
	USE ApiCoin = "USE"
	BCZ ApiCoin = "BCZ"
	IOC ApiCoin = "IOC"
	ASF ApiCoin = "ASF"
	MASS ApiCoin = "MASS"
	FAIR ApiCoin = "FAIR"
	NUKO ApiCoin = "NUKO"
	CMT ApiCoin = "CMT"
	EUNO ApiCoin = "EUNO"
	IOTX ApiCoin = "IOTX"
	ONION ApiCoin = "ONION"
	BTS ApiCoin = "BTS"
	UGAS ApiCoin = "UGAS"
	ADS ApiCoin = "ADS"
	ARA ApiCoin = "ARA"
	ZIL ApiCoin = "ZIL"
	MOAC ApiCoin = "MOAC"
	SWTC ApiCoin = "SWTC"
	VNSC ApiCoin = "VNSC"
	ECC ApiCoin = "ECC"
	RPD ApiCoin = "RPD"
	RAP ApiCoin = "RAP"
	GARD ApiCoin = "GARD"
	ZER ApiCoin = "ZER"
	EBST ApiCoin = "EBST"
	SHARD ApiCoin = "SHARD"
	CMM ApiCoin = "CMM"
	BLOCK ApiCoin = "BLOCK"
	AUDAX ApiCoin = "AUDAX"
	LUNA ApiCoin = "LUNA"
	ZPM ApiCoin = "ZPM"
	MEM ApiCoin = "MEM"
	SWIFT ApiCoin = "SWIFT"
	FIX ApiCoin = "FIX"
	CPC ApiCoin = "CPC"
	VGO ApiCoin = "VGO"
	DVT ApiCoin = "DVT"
	MTNS ApiCoin = "MTNS"
	BLAST ApiCoin = "BLAST"
	DCT ApiCoin = "DCT"
	AUX ApiCoin = "AUX"
	USDP ApiCoin = "USDP"
	HTDF ApiCoin = "HTDF"
	YEC ApiCoin = "YEC"
	ARW ApiCoin = "ARW"
	MDM ApiCoin = "MDM"
	CYB ApiCoin = "CYB"
	DOT ApiCoin = "DOT"
	AEON ApiCoin = "AEON"
	RES ApiCoin = "RES"
	AYA ApiCoin = "AYA"
	DAPS ApiCoin = "DAPS"
	CSC ApiCoin = "CSC"
	NOLLAR ApiCoin = "NOLLAR"
	XNOS ApiCoin = "XNOS"
	CPU ApiCoin = "CPU"
	VCT ApiCoin = "VCT"
	CZR ApiCoin = "CZR"
	ABBC ApiCoin = "ABBC"
	HET ApiCoin = "HET"
	XAS ApiCoin = "XAS"
	VDL ApiCoin = "VDL"
	MED ApiCoin = "MED"
	ZVC ApiCoin = "ZVC"
	VESTX ApiCoin = "VESTX"
	DBT ApiCoin = "DBT"
	SEOS ApiCoin = "SEOS"
	MXW ApiCoin = "MXW"
	ZNZ ApiCoin = "ZNZ"
	XCX ApiCoin = "XCX"
	SOX ApiCoin = "SOX"
	NYZO ApiCoin = "NYZO"
	ULC ApiCoin = "ULC"
	KAL ApiCoin = "KAL"
	XSN ApiCoin = "XSN"
	DOGEC ApiCoin = "DOGEC"
	QBC ApiCoin = "QBC"
	IMG ApiCoin = "IMG"
	QOS ApiCoin = "QOS"
	PKT ApiCoin = "PKT"
	LHD ApiCoin = "LHD"
	CENNZ ApiCoin = "CENNZ"
	UMBRU ApiCoin = "UMBRU"
	EVER ApiCoin = "EVER"
	XPC ApiCoin = "XPC"
	NIX ApiCoin = "NIX"
	UC ApiCoin = "UC"
	GALI ApiCoin = "GALI"
	OLT ApiCoin = "OLT"
	XBI ApiCoin = "XBI"
	DONU ApiCoin = "DONU"
	EARTHS ApiCoin = "EARTHS"
	HDD ApiCoin = "HDD"
	SUGAR ApiCoin = "SUGAR"
	AILE ApiCoin = "AILE"
	TENT ApiCoin = "TENT"
	AIN ApiCoin = "AIN"
	MSR ApiCoin = "MSR"
	SUMO ApiCoin = "SUMO"
	ETN ApiCoin = "ETN"
	BYTZ ApiCoin = "BYTZ"
	WOW ApiCoin = "WOW"
	XTNC ApiCoin = "XTNC"
	LTHN ApiCoin = "LTHN"
	NODE ApiCoin = "NODE"
	AGM ApiCoin = "AGM"
	TELOS ApiCoin = "TELOS"
	AION ApiCoin = "AION"
	KTV ApiCoin = "KTV"
	ZCR ApiCoin = "ZCR"
	ERG ApiCoin = "ERG"
	PESO ApiCoin = "PESO"
	XRPHD ApiCoin = "XRPHD"
	KSM ApiCoin = "KSM"
	PCN ApiCoin = "PCN"
	NCH ApiCoin = "NCH"
	ICU ApiCoin = "ICU"
	FNSA ApiCoin = "FNSA"
	AERGO ApiCoin = "AERGO"
	XTH ApiCoin = "XTH"
	LV ApiCoin = "LV"
	PHR ApiCoin = "PHR"
	VITAE ApiCoin = "VITAE"
	DIN ApiCoin = "DIN"
	SPL ApiCoin = "SPL"
	YCE ApiCoin = "YCE"
	XLR ApiCoin = "XLR"
	DGLD ApiCoin = "DGLD"
	XNS ApiCoin = "XNS"
	EM ApiCoin = "EM"
	SHN ApiCoin = "SHN"
	SEELE ApiCoin = "SEELE"
	AE ApiCoin = "AE"
	ODX ApiCoin = "ODX"
	KAVA ApiCoin = "KAVA"
	GLEEC ApiCoin = "GLEEC"
	FIL ApiCoin = "FIL"
	RUTA ApiCoin = "RUTA"
	CSDT ApiCoin = "CSDT"
	ETI ApiCoin = "ETI"
	ERE ApiCoin = "ERE"
	BTH ApiCoin = "BTH"
	MESG ApiCoin = "MESG"
	FIMK ApiCoin = "FIMK"
	AR ApiCoin = "AR"
	OGO ApiCoin = "OGO"
	RNG ApiCoin = "RNG"
	PEXA ApiCoin = "PEXA"
	MOON ApiCoin = "MOON"
	FCH ApiCoin = "FCH"
	LAT ApiCoin = "LAT"
	VEO ApiCoin = "VEO"
	GFN ApiCoin = "GFN"
	BAND ApiCoin = "BAND"
	DROP ApiCoin = "DROP"
	LYRA ApiCoin = "LYRA"
	CS ApiCoin = "CS"
	RUPX ApiCoin = "RUPX"
	THETA ApiCoin = "THETA"
	SOL ApiCoin = "SOL"
	THT ApiCoin = "THT"
	CFX ApiCoin = "CFX"
	KUMA ApiCoin = "KUMA"
	HASH ApiCoin = "HASH"
	CSPR ApiCoin = "CSPR"
	EARTH ApiCoin = "EARTH"
	EGLD ApiCoin = "EGLD"
	CHI ApiCoin = "CHI"
	KOTO ApiCoin = "KOTO"
	XRD ApiCoin = "XRD"
	AETH ApiCoin = "AETH"
	DNA ApiCoin = "DNA"
	SIERRA ApiCoin = "SIERRA"
	LET ApiCoin = "LET"
	BTCV ApiCoin = "BTCV"
	ABA ApiCoin = "ABA"
	SCC ApiCoin = "SCC"
	EDG ApiCoin = "EDG"
	AMS ApiCoin = "AMS"
	BU ApiCoin = "BU"
	GRAM ApiCoin = "GRAM"
	YAP ApiCoin = "YAP"
	NOVO ApiCoin = "NOVO"
	GHOST ApiCoin = "GHOST"
	HST ApiCoin = "HST"
	PRJ ApiCoin = "PRJ"
	YOU ApiCoin = "YOU"
	BYND ApiCoin = "BYND"
	FLOW ApiCoin = "FLOW"
	SCDO ApiCoin = "SCDO"
	BIND ApiCoin = "BIND"
	COINEVO ApiCoin = "COINEVO"
	SCRIBE ApiCoin = "SCRIBE"
	HYN ApiCoin = "HYN"
	BHP ApiCoin = "BHP"
	MKF ApiCoin = "MKF"
	XDC ApiCoin = "XDC"
	STR ApiCoin = "STR"
	HBC ApiCoin = "HBC"
	KTS ApiCoin = "KTS"
	LKR ApiCoin = "LKR"
	XWC ApiCoin = "XWC"
	DEAL ApiCoin = "DEAL"
	NTY ApiCoin = "NTY"
	AG ApiCoin = "AG"
	CICO ApiCoin = "CICO"
	IRIS ApiCoin = "IRIS"
	BDX ApiCoin = "BDX"
	SLS ApiCoin = "SLS"
	SRM ApiCoin = "SRM"
	BPS ApiCoin = "BPS"
	NKN ApiCoin = "NKN"
	ICL ApiCoin = "ICL"
	BONO ApiCoin = "BONO"
	PLC ApiCoin = "PLC"
	DUN ApiCoin = "DUN"
	DMCH ApiCoin = "DMCH"
	CTC ApiCoin = "CTC"
	GBCR ApiCoin = "GBCR"
	XDAG ApiCoin = "XDAG"
	SCAP ApiCoin = "SCAP"
	GTM ApiCoin = "GTM"
	RNL ApiCoin = "RNL"
	GRIN ApiCoin = "GRIN"
	MWC ApiCoin = "MWC"
	DOCK ApiCoin = "DOCK"
	POLYX ApiCoin = "POLYX"
	DIVER ApiCoin = "DIVER"
	APN ApiCoin = "APN"
	MTC ApiCoin = "MTC"
	NC ApiCoin = "NC"
	XINY ApiCoin = "XINY"
	BUFS ApiCoin = "BUFS"
	STOS ApiCoin = "STOS"
	TON ApiCoin = "TON"
	TAFT ApiCoin = "TAFT"
	HYDRA ApiCoin = "HYDRA"
	NOR ApiCoin = "NOR"
	WCN ApiCoin = "WCN"
	PSWAP ApiCoin = "PSWAP"
	XOR ApiCoin = "XOR"
	SSP ApiCoin = "SSP"
	DEI ApiCoin = "DEI"
	ZERO ApiCoin = "ZERO"
	ALPHA ApiCoin = "ALPHA"
	NOBL ApiCoin = "NOBL"
	EAST ApiCoin = "EAST"
	KDA ApiCoin = "KDA"
	SOUL ApiCoin = "SOUL"
	LORE ApiCoin = "LORE"
	FNR ApiCoin = "FNR"
	NEXUS ApiCoin = "NEXUS"
	QTZ ApiCoin = "QTZ"
	XMA ApiCoin = "XMA"
	CALL ApiCoin = "CALL"
	VAL ApiCoin = "VAL"
	EMIT ApiCoin = "EMIT"
	APTOS ApiCoin = "APTOS"
	ADON ApiCoin = "ADON"
	BTSG ApiCoin = "BTSG"
	LFC ApiCoin = "LFC"
	TREE ApiCoin = "TREE"
	LX ApiCoin = "LX"
	XLN ApiCoin = "XLN"
	ZRB ApiCoin = "ZRB"
	UCO ApiCoin = "UCO"
	WMP ApiCoin = "WMP"
	KOIN ApiCoin = "KOIN"
	PIRATE ApiCoin = "PIRATE"
	ACT ApiCoin = "ACT"
	PRKL ApiCoin = "PRKL"
	SSC ApiCoin = "SSC"
	GC ApiCoin = "GC"
	PLGR ApiCoin = "PLGR"
	MPLGR ApiCoin = "MPLGR"
	KNOX ApiCoin = "KNOX"
	ZED ApiCoin = "ZED"
	CNDL ApiCoin = "CNDL"
	WLKRR ApiCoin = "WLKRR"
	YUNGE ApiCoin = "YUNGE"
	VOKEN ApiCoin = "VOKEN"
	APL ApiCoin = "APL"
	EVRYNET ApiCoin = "EVRYNET"
	NENG ApiCoin = "NENG"
	CHTA ApiCoin = "CHTA"
	OAS ApiCoin = "OAS"
	KLV ApiCoin = "KLV"
	VEIL ApiCoin = "VEIL"
	GTB ApiCoin = "GTB"
	XDAI ApiCoin = "XDAI"
	COM ApiCoin = "COM"
	CHC ApiCoin = "CHC"
	SERF ApiCoin = "SERF"
	BNB ApiCoin = "BNB"
	SIN ApiCoin = "SIN"
	DLN ApiCoin = "DLN"
	BONTE ApiCoin = "BONTE"
	PEER ApiCoin = "PEER"
	ZET ApiCoin = "ZET"
	ABY ApiCoin = "ABY"
	XVC ApiCoin = "XVC"
	MCX ApiCoin = "MCX"
	HEALIOS ApiCoin = "HEALIOS"
	BMK ApiCoin = "BMK"
	DENTX ApiCoin = "DENTX"
	CFG ApiCoin = "CFG"
	XPRT ApiCoin = "XPRT"
	HONEY ApiCoin = "HONEY"
	BALLZ ApiCoin = "BALLZ"
	COSA ApiCoin = "COSA"
	BR ApiCoin = "BR"
	SUI ApiCoin = "SUI"
	UIDD ApiCoin = "UIDD"
	ACA ApiCoin = "ACA"
	BNC ApiCoin = "BNC"
	TAU ApiCoin = "TAU"
	PDEX ApiCoin = "PDEX"
	ZKS ApiCoin = "ZKS"
	QVT ApiCoin = "QVT"
	MEER ApiCoin = "MEER"
	REEF ApiCoin = "REEF"
	CLO ApiCoin = "CLO"
	BDB ApiCoin = "BDB"
	ACE ApiCoin = "ACE"
	CCN ApiCoin = "CCN"
	BBA ApiCoin = "BBA"
	CRUZ ApiCoin = "CRUZ"
	SAPP ApiCoin = "SAPP"
	KYAN ApiCoin = "KYAN"
	AZR ApiCoin = "AZR"
	CFL ApiCoin = "CFL"
	TRTT ApiCoin = "TRTT"
	PNY ApiCoin = "PNY"
	BECN ApiCoin = "BECN"
	MONK ApiCoin = "MONK"
	SAGA ApiCoin = "SAGA"
	SUV ApiCoin = "SUV"
	ESK ApiCoin = "ESK"
	BIR ApiCoin = "BIR"
	MOBIC ApiCoin = "MOBIC"
	FLS ApiCoin = "FLS"
	DSM ApiCoin = "DSM"
	HVH ApiCoin = "HVH"
	MOB ApiCoin = "MOB"
	IF ApiCoin = "IF"
	NAM ApiCoin = "NAM"
	ZBC ApiCoin = "ZBC"
	NEO ApiCoin = "NEO"
	TOMO ApiCoin = "TOMO"
	XSEL ApiCoin = "XSEL"
	LKSC ApiCoin = "LKSC"
	AS ApiCoin = "AS"
	XEC ApiCoin = "XEC"
	LMO ApiCoin = "LMO"
	HNT ApiCoin = "HNT"
	FIS ApiCoin = "FIS"
	SGE ApiCoin = "SGE"
	GERT ApiCoin = "GERT"
	META ApiCoin = "META"
	FRA ApiCoin = "FRA"
	CCD ApiCoin = "CCD"
	GHM ApiCoin = "GHM"
	LTP ApiCoin = "LTP"
	VKAX ApiCoin = "VKAX"
	MATIC ApiCoin = "MATIC"
	UNW ApiCoin = "UNW"
	TWINS ApiCoin = "TWINS"
	TLOS ApiCoin = "TLOS"
	AU ApiCoin = "AU"
	VCG ApiCoin = "VCG"
	AIOZ ApiCoin = "AIOZ"
	CORE ApiCoin = "CORE"
	PEC ApiCoin = "PEC"
	UNT ApiCoin = "UNT"
	CAPS ApiCoin = "CAPS"
	SUM ApiCoin = "SUM"
	TT ApiCoin = "TT"
	BKT ApiCoin = "BKT"
	NODL ApiCoin = "NODL"
	PCOIN ApiCoin = "PCOIN"
	TAO ApiCoin = "TAO"
	FTM ApiCoin = "FTM"
	RPG ApiCoin = "RPG"
	LAKE ApiCoin = "LAKE"
	ELV ApiCoin = "ELV"
	BIC ApiCoin = "BIC"
	EVC ApiCoin = "EVC"
	ONT ApiCoin = "ONT"
	CZZ ApiCoin = "CZZ"
	MCM ApiCoin = "MCM"
	BTCR ApiCoin = "BTCR"
	RISE ApiCoin = "RISE"
	DFI ApiCoin = "DFI"
	EFI ApiCoin = "EFI"
	ALPH ApiCoin = "ALPH"
	GLMR ApiCoin = "GLMR"
	MOVR ApiCoin = "MOVR"
	WPC ApiCoin = "WPC"
	WEI ApiCoin = "WEI"
	DFC ApiCoin = "DFC"
	HYC ApiCoin = "HYC"
	BEAM ApiCoin = "BEAM"
	ELF ApiCoin = "ELF"
	AUDL ApiCoin = "AUDL"
	ATH ApiCoin = "ATH"
	NEW ApiCoin = "NEW"
	BTA ApiCoin = "BTA"
	NEOX ApiCoin = "NEOX"
	MEWC ApiCoin = "MEWC"
	BCX ApiCoin = "BCX"
	XTZ ApiCoin = "XTZ"
	BBP ApiCoin = "BBP"
	ADA ApiCoin = "ADA"
	TES ApiCoin = "TES"
	ZTX ApiCoin = "ZTX"
	CLC ApiCoin = "CLC"
	VIPS ApiCoin = "VIPS"
	XMX ApiCoin = "XMX"
	TRTL ApiCoin = "TRTL"
	EGEM ApiCoin = "EGEM"
	HODL ApiCoin = "HODL"
	PHL ApiCoin = "PHL"
	SC ApiCoin = "SC"
	MYT ApiCoin = "MYT"
	POLIS ApiCoin = "POLIS"
	XMCC ApiCoin = "XMCC"
	COLX ApiCoin = "COLX"
	GIN ApiCoin = "GIN"
	MNP ApiCoin = "MNP"
	MLN ApiCoin = "MLN"
	KIN ApiCoin = "KIN"
	EOSC ApiCoin = "EOSC"
	PKC ApiCoin = "PKC"
	SKT ApiCoin = "SKT"
	ANY ApiCoin = "ANY"
	MCASH ApiCoin = "MCASH"
	TRUE ApiCoin = "TRUE"
	IOTE ApiCoin = "IOTE"
	BAY ApiCoin = "BAY"
	XRG ApiCoin = "XRG"
	CHZ ApiCoin = "CHZ"
	ASK ApiCoin = "ASK"
	QTUM ApiCoin = "QTUM"
	ETP ApiCoin = "ETP"
	GXC ApiCoin = "GXC"
	CRP ApiCoin = "CRP"
	ELA ApiCoin = "ELA"
	SNOW ApiCoin = "SNOW"
	XIN ApiCoin = "XIN"
	NEXI ApiCoin = "NEXI"
	AOA ApiCoin = "AOA"
	NAS ApiCoin = "NAS"
	BND ApiCoin = "BND"
	LUX ApiCoin = "LUX"
	COS ApiCoin = "COS"
	CCC ApiCoin = "CCC"
	SXP ApiCoin = "SXP"
	ROI ApiCoin = "ROI"
	DYN ApiCoin = "DYN"
	SEQ ApiCoin = "SEQ"
	DEO ApiCoin = "DEO"
	DST ApiCoin = "DST"
	CY ApiCoin = "CY"
	YEE ApiCoin = "YEE"
	IOTA ApiCoin = "IOTA"
	SMR ApiCoin = "SMR"
	AXE ApiCoin = "AXE"
	XYM ApiCoin = "XYM"
	XVM ApiCoin = "XVM"
	FIC ApiCoin = "FIC"
	HNS ApiCoin = "HNS"
	ISK ApiCoin = "ISK"
	ALTME ApiCoin = "ALTME"
	FUND ApiCoin = "FUND"
	STX ApiCoin = "STX"
	SLU ApiCoin = "SLU"
	COTI ApiCoin = "COTI"
	ROGER ApiCoin = "ROGER"
	TOPL ApiCoin = "TOPL"
	KLY ApiCoin = "KLY"
	SHFT ApiCoin = "SHFT"
	BTV ApiCoin = "BTV"
	SKY ApiCoin = "SKY"
	KLAY ApiCoin = "KLAY"
	BTQ ApiCoin = "BTQ"
	XCH ApiCoin = "XCH"
	PLMNT ApiCoin = "PLMNT"
	NULS ApiCoin = "NULS"
	BBC ApiCoin = "BBC"
	JGC ApiCoin = "JGC"
	AVAX ApiCoin = "AVAX"
	BOBA ApiCoin = "BOBA"
	LOOP ApiCoin = "LOOP"
	STRK ApiCoin = "STRK"
	NRG ApiCoin = "NRG"
	FO ApiCoin = "FO"
	RTM ApiCoin = "RTM"
	XRC ApiCoin = "XRC"
	XPI ApiCoin = "XPI"
	VARCH ApiCoin = "VARCH"
	TNKR ApiCoin = "TNKR"
	IPOS ApiCoin = "IPOS"
	MINA ApiCoin = "MINA"
	BTY ApiCoin = "BTY"
	SDGO ApiCoin = "SDGO"
	ARDR ApiCoin = "ARDR"
	MTR ApiCoin = "MTR"
	SAFE ApiCoin = "SAFE"
	FLUX ApiCoin = "FLUX"
	RITO ApiCoin = "RITO"
	XND ApiCoin = "XND"
	PAC ApiCoin = "PAC"
	PWR ApiCoin = "PWR"
	BELL ApiCoin = "BELL"
	CHX ApiCoin = "CHX"
	NEXA ApiCoin = "NEXA"
	BTT ApiCoin = "BTT"
	FXTC ApiCoin = "FXTC"
	AMA ApiCoin = "AMA"
	AXIV ApiCoin = "AXIV"
	EVE ApiCoin = "EVE"
	STASH ApiCoin = "STASH"
	CELO ApiCoin = "CELO"
	TH ApiCoin = "TH"
	GRLC ApiCoin = "GRLC"
	GWL ApiCoin = "GWL"
	ZYN ApiCoin = "ZYN"
	WICC ApiCoin = "WICC"
	HOME ApiCoin = "HOME"
	STC ApiCoin = "STC"
	STRAX ApiCoin = "STRAX"
	KAS ApiCoin = "KAS"
	AKA ApiCoin = "AKA"
	GENOM ApiCoin = "GENOM"
	ZAMA ApiCoin = "ZAMA"
	SCR ApiCoin = "SCR"
	VITE ApiCoin = "VITE"
	WTC ApiCoin = "WTC"
	ILT ApiCoin = "ILT"
	XERO ApiCoin = "XERO"
	LAX ApiCoin = "LAX"
	BCO ApiCoin = "BCO"
	BHD ApiCoin = "BHD"
	PTN ApiCoin = "PTN"
	VLX ApiCoin = "VLX"
	WAN ApiCoin = "WAN"
	WAVES ApiCoin = "WAVES"
	ABC ApiCoin = "ABC"
	CRM ApiCoin = "CRM"
	SEM ApiCoin = "SEM"
	ION ApiCoin = "ION"
	FCT ApiCoin = "FCT"
	WGR ApiCoin = "WGR"
	OBSR ApiCoin = "OBSR"
	AFS ApiCoin = "AFS"
	XDS ApiCoin = "XDS"
	AQUA ApiCoin = "AQUA"
	HATCH ApiCoin = "HATCH"
	KUSD ApiCoin = "KUSD"
	GENS ApiCoin = "GENS"
	EQ ApiCoin = "EQ"
	QKC ApiCoin = "QKC"
	FVDC ApiCoin = "FVDC"
	USDC ApiCoin = "USDC"
)

// All allowed values of ApiCoin enum
var AllowedApiCoinEnumValues = []ApiCoin{
	"BTC",
	"LTC",
	"DOGE",
	"RDD",
	"DASH",
	"PPC",
	"NMC",
	"FTC",
	"XCP",
	"BLK",
	"NSR",
	"NBT",
	"MZC",
	"VIA",
	"RBY",
	"GRS",
	"DGC",
	"DGB",
	"MONA",
	"CLAM",
	"XPM",
	"NEOS",
	"JBS",
	"ZRC",
	"VTC",
	"NXT",
	"BURST",
	"MUE",
	"ZOOM",
	"SDC",
	"PKB",
	"PND",
	"START",
	"MOIN",
	"EXP",
	"DCR",
	"XEM",
	"PART",
	"SHR",
	"NVC",
	"AC",
	"BTCD",
	"DOPE",
	"TPC",
	"AIB",
	"EDRC",
	"SYS",
	"SLR",
	"SMLY",
	"ETH",
	"PSB",
	"XBC",
	"NXS",
	"INSN",
	"OK",
	"BRIT",
	"CMP",
	"CRW",
	"BELA",
	"ICX",
	"FJC",
	"MIX",
	"CLUB",
	"RICHX",
	"POT",
	"QRK",
	"TRC",
	"GRC",
	"AUR",
	"IXC",
	"NLG",
	"BITB",
	"XMY",
	"BSD",
	"UNO",
	"GB",
	"SHM",
	"CRX",
	"BIQ",
	"EVO",
	"STO",
	"BIGUP",
	"GAME",
	"DLC",
	"ZYD",
	"DBIC",
	"STRAT",
	"SH",
	"MARS",
	"UBQ",
	"PTC",
	"NRO",
	"ARK",
	"USC",
	"THC",
	"LINX",
	"ECN",
	"DNR",
	"PINK",
	"ATOM",
	"PIVX",
	"FLASH",
	"ZEN",
	"PUT",
	"ZNY",
	"UNIFY",
	"XST",
	"VC",
	"XMR",
	"VOX",
	"NAV",
	"ZEC",
	"LSK",
	"STEEM",
	"XZC",
	"RBTC",
	"RPT",
	"KMD",
	"RIC",
	"XRP",
	"NEBL",
	"ZCL",
	"WHL",
	"ERC",
	"DMD",
	"BTM",
	"BIO",
	"SSN",
	"TOA",
	"BTX",
	"ACC",
	"ELLA",
	"PIRL",
	"XNO",
	"VIVO",
	"FRST",
	"HNC",
	"BUZZ",
	"MBRS",
	"HC",
	"HTML",
	"ODN",
	"ONX",
	"RVN",
	"GBX",
	"BTCZ",
	"POA",
	"NYC",
	"MXT",
	"WC",
	"MNX",
	"MUSIC",
	"CRAVE",
	"STAK",
	"LCH",
	"EXCL",
	"LCC",
	"XFE",
	"EOS",
	"TRX",
	"KOBO",
	"HUSH",
	"BAN",
	"ETF",
	"OMNI",
	"BIFI",
	"CNMC",
	"BCN",
	"RIN",
	"ATP",
	"EVT",
	"ATN",
	"BIS",
	"NEET",
	"BOPO",
	"OOT",
	"ALIAS",
	"BOXY",
	"FLO",
	"MEC",
	"BTDX",
	"XAX",
	"ANON",
	"LTZ",
	"SMART",
	"XUEZ",
	"HLM",
	"WEB",
	"ACM",
	"BITC",
	"TZC",
	"VAR",
	"IOV",
	"FIO",
	"BSV",
	"DXN",
	"PCX",
	"LOKI",
	"NIM",
	"EXOS",
	"ECA",
	"SOOM",
	"FREE",
	"NPW",
	"BST",
	"ZEST",
	"ABT",
	"PION",
	"ZBUX",
	"KPL",
	"TPAY",
	"ZILLA",
	"ANK",
	"BCC",
	"HPB",
	"ONE",
	"SBC",
	"IPC",
	"DMTC",
	"OGC",
	"SHIT",
	"ANDES",
	"AREPA",
	"BOLI",
	"RIL",
	"BRAVO",
	"ALGO",
	"BZX",
	"GXX",
	"HEAT",
	"XDN",
	"FSN",
	"BOLD",
	"IOST",
	"TKEY",
	"USE",
	"BCZ",
	"IOC",
	"ASF",
	"MASS",
	"FAIR",
	"NUKO",
	"CMT",
	"EUNO",
	"IOTX",
	"ONION",
	"BTS",
	"UGAS",
	"ADS",
	"ARA",
	"ZIL",
	"MOAC",
	"SWTC",
	"VNSC",
	"ECC",
	"RPD",
	"RAP",
	"GARD",
	"ZER",
	"EBST",
	"SHARD",
	"CMM",
	"BLOCK",
	"AUDAX",
	"LUNA",
	"ZPM",
	"MEM",
	"SWIFT",
	"FIX",
	"CPC",
	"VGO",
	"DVT",
	"MTNS",
	"BLAST",
	"DCT",
	"AUX",
	"USDP",
	"HTDF",
	"YEC",
	"ARW",
	"MDM",
	"CYB",
	"DOT",
	"AEON",
	"RES",
	"AYA",
	"DAPS",
	"CSC",
	"NOLLAR",
	"XNOS",
	"CPU",
	"VCT",
	"CZR",
	"ABBC",
	"HET",
	"XAS",
	"VDL",
	"MED",
	"ZVC",
	"VESTX",
	"DBT",
	"SEOS",
	"MXW",
	"ZNZ",
	"XCX",
	"SOX",
	"NYZO",
	"ULC",
	"KAL",
	"XSN",
	"DOGEC",
	"QBC",
	"IMG",
	"QOS",
	"PKT",
	"LHD",
	"CENNZ",
	"UMBRU",
	"EVER",
	"XPC",
	"NIX",
	"UC",
	"GALI",
	"OLT",
	"XBI",
	"DONU",
	"EARTHS",
	"HDD",
	"SUGAR",
	"AILE",
	"TENT",
	"AIN",
	"MSR",
	"SUMO",
	"ETN",
	"BYTZ",
	"WOW",
	"XTNC",
	"LTHN",
	"NODE",
	"AGM",
	"TELOS",
	"AION",
	"KTV",
	"ZCR",
	"ERG",
	"PESO",
	"XRPHD",
	"KSM",
	"PCN",
	"NCH",
	"ICU",
	"FNSA",
	"AERGO",
	"XTH",
	"LV",
	"PHR",
	"VITAE",
	"DIN",
	"SPL",
	"YCE",
	"XLR",
	"DGLD",
	"XNS",
	"EM",
	"SHN",
	"SEELE",
	"AE",
	"ODX",
	"KAVA",
	"GLEEC",
	"FIL",
	"RUTA",
	"CSDT",
	"ETI",
	"ERE",
	"BTH",
	"MESG",
	"FIMK",
	"AR",
	"OGO",
	"RNG",
	"PEXA",
	"MOON",
	"FCH",
	"LAT",
	"VEO",
	"GFN",
	"BAND",
	"DROP",
	"LYRA",
	"CS",
	"RUPX",
	"THETA",
	"SOL",
	"THT",
	"CFX",
	"KUMA",
	"HASH",
	"CSPR",
	"EARTH",
	"EGLD",
	"CHI",
	"KOTO",
	"XRD",
	"AETH",
	"DNA",
	"SIERRA",
	"LET",
	"BTCV",
	"ABA",
	"SCC",
	"EDG",
	"AMS",
	"BU",
	"GRAM",
	"YAP",
	"NOVO",
	"GHOST",
	"HST",
	"PRJ",
	"YOU",
	"BYND",
	"FLOW",
	"SCDO",
	"BIND",
	"COINEVO",
	"SCRIBE",
	"HYN",
	"BHP",
	"MKF",
	"XDC",
	"STR",
	"HBC",
	"KTS",
	"LKR",
	"XWC",
	"DEAL",
	"NTY",
	"AG",
	"CICO",
	"IRIS",
	"BDX",
	"SLS",
	"SRM",
	"BPS",
	"NKN",
	"ICL",
	"BONO",
	"PLC",
	"DUN",
	"DMCH",
	"CTC",
	"GBCR",
	"XDAG",
	"SCAP",
	"GTM",
	"RNL",
	"GRIN",
	"MWC",
	"DOCK",
	"POLYX",
	"DIVER",
	"APN",
	"MTC",
	"NC",
	"XINY",
	"BUFS",
	"STOS",
	"TON",
	"TAFT",
	"HYDRA",
	"NOR",
	"WCN",
	"PSWAP",
	"XOR",
	"SSP",
	"DEI",
	"ZERO",
	"ALPHA",
	"NOBL",
	"EAST",
	"KDA",
	"SOUL",
	"LORE",
	"FNR",
	"NEXUS",
	"QTZ",
	"XMA",
	"CALL",
	"VAL",
	"EMIT",
	"APTOS",
	"ADON",
	"BTSG",
	"LFC",
	"TREE",
	"LX",
	"XLN",
	"ZRB",
	"UCO",
	"WMP",
	"KOIN",
	"PIRATE",
	"ACT",
	"PRKL",
	"SSC",
	"GC",
	"PLGR",
	"MPLGR",
	"KNOX",
	"ZED",
	"CNDL",
	"WLKRR",
	"YUNGE",
	"VOKEN",
	"APL",
	"EVRYNET",
	"NENG",
	"CHTA",
	"OAS",
	"KLV",
	"VEIL",
	"GTB",
	"XDAI",
	"COM",
	"CHC",
	"SERF",
	"BNB",
	"SIN",
	"DLN",
	"BONTE",
	"PEER",
	"ZET",
	"ABY",
	"XVC",
	"MCX",
	"HEALIOS",
	"BMK",
	"DENTX",
	"CFG",
	"XPRT",
	"HONEY",
	"BALLZ",
	"COSA",
	"BR",
	"SUI",
	"UIDD",
	"ACA",
	"BNC",
	"TAU",
	"PDEX",
	"ZKS",
	"QVT",
	"MEER",
	"REEF",
	"CLO",
	"BDB",
	"ACE",
	"CCN",
	"BBA",
	"CRUZ",
	"SAPP",
	"KYAN",
	"AZR",
	"CFL",
	"TRTT",
	"PNY",
	"BECN",
	"MONK",
	"SAGA",
	"SUV",
	"ESK",
	"BIR",
	"MOBIC",
	"FLS",
	"DSM",
	"HVH",
	"MOB",
	"IF",
	"NAM",
	"ZBC",
	"NEO",
	"TOMO",
	"XSEL",
	"LKSC",
	"AS",
	"XEC",
	"LMO",
	"HNT",
	"FIS",
	"SGE",
	"GERT",
	"META",
	"FRA",
	"CCD",
	"GHM",
	"LTP",
	"VKAX",
	"MATIC",
	"UNW",
	"TWINS",
	"TLOS",
	"AU",
	"VCG",
	"AIOZ",
	"CORE",
	"PEC",
	"UNT",
	"CAPS",
	"SUM",
	"TT",
	"BKT",
	"NODL",
	"PCOIN",
	"TAO",
	"FTM",
	"RPG",
	"LAKE",
	"ELV",
	"BIC",
	"EVC",
	"ONT",
	"CZZ",
	"MCM",
	"BTCR",
	"RISE",
	"DFI",
	"EFI",
	"ALPH",
	"GLMR",
	"MOVR",
	"WPC",
	"WEI",
	"DFC",
	"HYC",
	"BEAM",
	"ELF",
	"AUDL",
	"ATH",
	"NEW",
	"BTA",
	"NEOX",
	"MEWC",
	"BCX",
	"XTZ",
	"BBP",
	"ADA",
	"TES",
	"ZTX",
	"CLC",
	"VIPS",
	"XMX",
	"TRTL",
	"EGEM",
	"HODL",
	"PHL",
	"SC",
	"MYT",
	"POLIS",
	"XMCC",
	"COLX",
	"GIN",
	"MNP",
	"MLN",
	"KIN",
	"EOSC",
	"PKC",
	"SKT",
	"ANY",
	"MCASH",
	"TRUE",
	"IOTE",
	"BAY",
	"XRG",
	"CHZ",
	"ASK",
	"QTUM",
	"ETP",
	"GXC",
	"CRP",
	"ELA",
	"SNOW",
	"XIN",
	"NEXI",
	"AOA",
	"NAS",
	"BND",
	"LUX",
	"COS",
	"CCC",
	"SXP",
	"ROI",
	"DYN",
	"SEQ",
	"DEO",
	"DST",
	"CY",
	"YEE",
	"IOTA",
	"SMR",
	"AXE",
	"XYM",
	"XVM",
	"FIC",
	"HNS",
	"ISK",
	"ALTME",
	"FUND",
	"STX",
	"SLU",
	"COTI",
	"ROGER",
	"TOPL",
	"KLY",
	"SHFT",
	"BTV",
	"SKY",
	"KLAY",
	"BTQ",
	"XCH",
	"PLMNT",
	"NULS",
	"BBC",
	"JGC",
	"AVAX",
	"BOBA",
	"LOOP",
	"STRK",
	"NRG",
	"FO",
	"RTM",
	"XRC",
	"XPI",
	"VARCH",
	"TNKR",
	"IPOS",
	"MINA",
	"BTY",
	"SDGO",
	"ARDR",
	"MTR",
	"SAFE",
	"FLUX",
	"RITO",
	"XND",
	"PAC",
	"PWR",
	"BELL",
	"CHX",
	"NEXA",
	"BTT",
	"FXTC",
	"AMA",
	"AXIV",
	"EVE",
	"STASH",
	"CELO",
	"TH",
	"GRLC",
	"GWL",
	"ZYN",
	"WICC",
	"HOME",
	"STC",
	"STRAX",
	"KAS",
	"AKA",
	"GENOM",
	"ZAMA",
	"SCR",
	"VITE",
	"WTC",
	"ILT",
	"XERO",
	"LAX",
	"BCO",
	"BHD",
	"PTN",
	"VLX",
	"WAN",
	"WAVES",
	"ABC",
	"CRM",
	"SEM",
	"ION",
	"FCT",
	"WGR",
	"OBSR",
	"AFS",
	"XDS",
	"AQUA",
	"HATCH",
	"KUSD",
	"GENS",
	"EQ",
	"QKC",
	"FVDC",
	"USDC",
}

func (v *ApiCoin) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiCoin(value)
	for _, existing := range AllowedApiCoinEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiCoin", value)
}

// NewApiCoinFromValue returns a pointer to a valid ApiCoin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiCoinFromValue(v string) (*ApiCoin, error) {
	ev := ApiCoin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiCoin: valid values are %v", v, AllowedApiCoinEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiCoin) IsValid() bool {
	for _, existing := range AllowedApiCoinEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to api.Coin value
func (v ApiCoin) Ptr() *ApiCoin {
	return &v
}

type NullableApiCoin struct {
	value *ApiCoin
	isSet bool
}

func (v NullableApiCoin) Get() *ApiCoin {
	return v.value
}

func (v *NullableApiCoin) Set(val *ApiCoin) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCoin) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCoin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCoin(val *ApiCoin) *NullableApiCoin {
	return &NullableApiCoin{value: val, isSet: true}
}

func (v NullableApiCoin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCoin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

