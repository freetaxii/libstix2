// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package vocabs

// Account - This defines the STIX account vocabulary.
var Account = []string{
	"facebook",
	"ldap",
	"nis",
	"openid",
	"radius",
	"skype",
	"tacacs",
	"twitter",
	"unix",
	"windows-local",
	"windows-domain",
}

// AttackMotivation - This defines the STIX attack motivation vocabulary.
var AttackMotivation = []string{
	"accidental",
	"coercion",
	"dominance",
	"ideology",
	"notoriety",
	"organizational-gain",
	"personal-gain",
	"personal-satisfaction",
	"revenge",
	"unpredictable",
}

// AttackResourceLevel - This defines the STIX attack resource level vocabulary.
var AttackResourceLevel = []string{
	"individual",
	"club",
	"contest",
	"team",
	"organization",
	"government",
}

// CourseOfAction - This defines the STIX course of action vocabulary.
var CourseOfAction = []string{
	"textual:text/plain",
	"textual:text/html",
	"textual:text/md",
	"textual:pdf",
}

// Encryption - This defines the STIX encryption enumeration.
var Encryption = []string{
	"AES-256-GCM",
	"ChaCha20-Poly1305",
	"mime-type-indicated",
}

// Grouping - This defines the STIX grouping vocabulary.
var Grouping = []string{
	"suspicious-activity",
	"malware-analysis",
	"unspecified",
}

// HashingAlgorithm - This defines the STIX hashing algorithm vocabulary.
var HashingAlgorithm = []string{
	"MD5",
	"SHA-1",
	"SHA-256",
	"SHA-512",
	"SHA3-256",
	"SHA3-512",
	"SSDEEP",
	"TLSH",
}

// IdentityClass - This defines the STIX identity class vocabulary.
var IdentityClass = []string{
	"individual",
	"group",
	"system",
	"organization",
	"class",
	"unspecified",
}

// IndicatorLabel - This defines the STIX indicator label vocabulary.
var IndicatorLabel = []string{
	"anomalous-activity",
	"anonymization",
	"benign",
	"compromised",
	"malicious-activity",
	"attribution",
}

// ImplementationLanguage - This defines the STIX implementation language vocabulary.
var ImplementationLanguage = []string{
	"applescript",
	"bash",
	"c",
	"c++",
	"c#",
	"go",
	"java",
	"javascript",
	"lua",
	"objective-c",
	"perl",
	"php",
	"powershell",
	"python",
	"ruby",
	"scala",
	"swift",
	"typescript",
	"visual-basic",
	"x86-32",
	"x86-64",
}

// IndicatorType - This defines the STIX indicator type vocabulary.
var IndicatorType = []string{
	"anomalous-activity",
	"anonymization",
	"benign",
	"compromised",
	"malicious-activity",
	"attribution",
	"unknown",
}

// IndustrySector - This defines the STIX industry sector vocabulary.
var IndustrySector = []string{
	"agriculture",
	"aerospace",
	"automotive",
	"communications",
	"construction",
	"defense",
	"education",
	"energy",
	"entertainment",
	"financial-services",
	"government-national",
	"government-regional",
	"government-local",
	"government-public-services",
	"healthcare",
	"hospitality-leisure",
	"infrastructure",
	"insurance",
	"manufacturing",
	"mining",
	"non-profit",
	"pharmaceuticals",
	"retail",
	"technology",
	"telecommunications",
	"transportation",
	"utilities",
}

// InfrastructureType - This defines the STIX infrastructure type vocabulary.
var InfrastructureType = []string{
	"amplification",
	"anonymization",
	"botnet",
	"command-and-control",
	"exfiltration",
	"hosting-malware",
	"hosting-target-lists",
	"phishing",
	"reconnaissance",
	"staging",
	"undefined",
}

// MalwareAVResults - This defines the STIX malware av results vocabulary.
var MalwareAVResults = []string{
	"malicious",
	"suspicious",
	"benign",
	"unknown",
}

// MalwareCapabilities - This defines the STIX malware capabilities vocabulary.
var MalwareCapabilities = []string{
	"accesses-remote-machines",
	"anti-debugging",
	"anti-disassembly",
	"anti-emulation",
	"anti-memory-forensics",
	"anti-sandbox",
	"anti-vm",
	"captures-input-peripherals",
	"captures-output-peripherals",
	"captures-system-state-data",
	"cleans-traces-of-infection",
	"commits-fraud",
	"communicates-with-c2",
	"compromises-data-availability",
	"compromises-data-integrity",
	"compromises-system-availability",
	"controls-local-machine",
	"degrades-security-software",
	"degrades-system-updates",
	"determines-c2-server",
	"emails-spam",
	"escalates-privileges",
	"evades-av",
	"exfiltrates-data",
	"fingerprints-host",
	"hides-artifacts",
	"hides-executing-code",
	"infects-files",
	"infects-remote-machines",
	"installs-other-components",
	"persists-after-system-reboot",
	"prevents-artifact-access",
	"prevents-artifact-deletion",
	"probes-network-environment",
	"self-modifies",
	"steals-authentication-credentials",
	"violates-system-operational-integrity",
}

// MalwareType - This defines the STIX malware type vocabulary.
var MalwareType = []string{
	"adware",
	"backdoor",
	"bot",
	"bootkit",
	"ddos",
	"downloader",
	"dropper",
	"exploit-kit",
	"keylogger",
	"ransomware",
	"remote-access-trojan",
	"resource-exploitation",
	"rogue-security-software",
	"rootkit",
	"screen-capture",
	"spyware",
	"trojan",
	"unknown",
	"virus",
	"webshell",
	"wiper",
	"worm",
}

// NetworkSocketAddressFamily - This defines the STIX network socket address family enumeration.
var NetworkSocketAddressFamily = []string{
	"AF_UNSPEC",
	"AF_INET",
	"AF_IPX",
	"AF_APPLETALK",
	"AF_NETBIOS",
	"AF_INET6",
	"AF_IRDA",
	"AF_BTH",
}

// NetworkSocketType - This defines the STIX network socket type enumeration.
var NetworkSocketType = []string{
	"SOCK_STREAM",
	"AF_ISOCK_DGRAMNET",
	"SOCK_RAW, SOCK_RDM",
	"SOCK_SEQPACKET",
}

// Opinion - This defines the STIX opinion enumeration.
var Opinion = []string{
	"strongly-disagree",
	"disagree",
	"neutral",
	"agree",
	"strongly-agree",
}

// PatternType - This defines the STIX pattern type vocabulary.
var PatternType = []string{
	"stix",
	"pcre",
	"sigma",
	"snort",
	"suricata",
	"yara",
}

// ProcessorArchitecture - This defines the STIX processor architecture vocabulary.
var ProcessorArchitecture = []string{
	"alpha",
	"arm",
	"ia-64",
	"mips",
	"powerpc",
	"sparc",
	"x86",
	"x86-64",
}

// Region - This defines the STIX region vocabulary.
var Region = []string{
	"africa",
	"eastern-africa",
	"middle-africa",
	"northern-africa",
	"southern-africa",
	"western-africa",
	"americas",
	"latin-america-caribbean",
	"south-america",
	"caribbean",
	"central-america",
	"northern-america",
	"asia",
	"central-asia",
	"eastern-asia",
	"southern-asia",
	"south-eastern-asia",
	"western-asia",
	"europe",
	"eastern-europe",
	"northern-europe",
	"southern-europe",
	"western-europe",
	"oceania",
	"antarctica",
	"australia-new-zealand",
	"melanesia",
	"micronesia",
	"polynesia",
}

// ReportType - This defines the STIX report type vocabulary.
var ReportType = []string{
	"attack-pattern",
	"campaign",
	"identity",
	"indicator",
	"intrusion-set",
	"malware",
	"observed-data",
	"threat-actor",
	"threat-report",
	"tool",
	"vulnerability",
}

// ThreatActorType - This defines the STIX threat actor type vocabulary.
var ThreatActorType = []string{
	"activist",
	"competitor",
	"crime-syndicate",
	"criminal",
	"hacker",
	"insider-accidental",
	"insider-disgruntled",
	"nation-state",
	"sensationalist",
	"spy",
	"terrorist",
	"unknown",
}

// ThreatActorRole - This defines the STIX threat actor role vocabulary.
var ThreatActorRole = []string{
	"agent",
	"director",
	"independent",
	"infrastructure-architect",
	"infrastructure-operator",
	"malware-author",
	"sponsor",
}

// ThreatActorSophistication - This defines the STIX threat actor sophistication vocabulary.
var ThreatActorSophistication = []string{
	"none",
	"minimal",
	"intermediate",
	"advanced",
	"expert",
	"innovator",
	"strategic",
}

// ToolType - This defines the STIX tool type vocabulary.
var ToolType = []string{
	"denial-of-service",
	"exploitation",
	"information-gathering",
	"network-capture",
	"credential-exploitation",
	"remote-access",
	"vulnerability-scanning",
	"unknown",
}

// WindowsIntegrityLevel - This defines the STIX Windows integrity level enumeration
var WindowsIntegrityLevel = []string{
	"low",
	"medium",
	"high",
	"system",
}

// WindowsPEBinary - This defines the STIX Windows pe binary vocabulary
var WindowsPEBinary = []string{
	"dll",
	"exe",
	"sys",
}

// WindowsRegistryDatatype - This defines the STIX Windows registry datatype enumeration.
var WindowsRegistryDatatype = []string{
	"REG_NONE",
	"REG_SZ",
	"REG_EXPAND_SZ",
	"REG_BINARY",
	"REG_DWORD",
	"REG_DWORD_BIG_ENDIAN",
	"REG_DWORD_LITTLE_ENDIAN",
	"REG_LINK",
	"REG_MULTI_SZ",
	"REG_RESOURCE_LIST",
	"REG_FULL_RESOURCE_DESCRIPTION",
	"REG_RESOURCE_REQUIREMENTS_LIST",
	"REG_QWORD",
	"REG_INVALID_TYPE",
}

// WindowsServiceStartType - This defines the STIX Windows service start type enumeration.
var WindowsServiceStartType = []string{
	"SERVICE_AUTO_START",
	"SERVICE_BOOT_START",
	"SERVICE_DEMAND_START",
	"SERVICE_DISABLED",
	"SERVICE_SYSTEM_ALERT",
}

// WindowsServiceType - This defines the STIX Windows service type enumeration.
var WindowsServiceType = []string{
	"SERVICE_KERNEL_DRIVER",
	"SERVICE_FILE_SYSTEM_DRIVER",
	"SERVICE_WIN32_OWN_PROCESS",
	"SERVICE_WIN32_SHARE_PROCESS",
}

// WindowsServiceStatus - This defines the STIX Windows service status enumeration.
var WindowsServiceStatus = []string{
	"SERVICE_CONTINUE_PENDING",
	"SERVICE_PAUSE_PENDING",
	"SERVICE_PAUSED",
	"SERVICE_RUNNING",
	"SERVICE_START_PENDING",
	"SERVICE_STOP_PENDING",
	"SERVICE_STOPPED",
}
