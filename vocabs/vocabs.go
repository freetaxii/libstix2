// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package vocabs

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
	"content",
	"team",
	"organization",
	"government",
}

// IdentityClass - This defines the STIX identity class vocabulary.
var IdentityClass = []string{
	"individual",
	"group",
	"organization",
	"class",
	"unknown",
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

// MalwareLabel - This defines the STIX malware label vocabulary.
var MalwareLabel = []string{
	"adware",
	"backdoor",
	"bot",
	"ddos",
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
	"virus",
	"worm",
}

// MalwareTypes - This defines the STIX malware types vocabulary.
var MalwareTypes = []string{
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

// ArchitectureExecutionEnvs This defines the STIX malware architecture execution envs vocabulary.
var ArchitectureExecutionEnvs = []string{
	"alpha",
	"arm",
	"ia-64",
	"mips",
	"powerpc",
	"sparc",
	"x86",
	"x86-64",
}

// ImplementationLanguages This defines the STIX malware implementation languages vocabulary.
var ImplementationLanguages = []string{
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

// MalwareCapabilities This defines the STIX malware malware capabilities vocabulary.
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

// InfrastructureTypes
var InfrastructureTypes = []string{
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

// ReportLabel - This defines the STIX report label vocabulary.
var ReportLabel = []string{
	"threat-report",
	"attack-pattern",
	"campaign",
	"indicator",
	"malware",
	"observed-data",
	"threat-actor",
	"tool",
	"victim-target",
	"vulnerability",
}

// ThreatActorLabel - This defines the STIX threat actor label vocabulary.
var ThreatActorLabel = []string{
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

// ToolLabel - This defines the STIX tool label vocabulary.
var ToolLabel = []string{
	"denial-of-service",
	"exploitation",
	"information-gathering",
	"network-capture",
	"credential-exploitation",
	"remote-access",
	"vulnerability-scanning",
}
