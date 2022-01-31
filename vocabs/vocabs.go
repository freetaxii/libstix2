// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package vocabs

// IsVocabEntryValid - This function determine will take in a vocabulary and a
// value and return true if it is found and false if it is not found.
func IsVocabEntryValid(vocab map[string]bool, s string) bool {
	if _, found := vocab[s]; found == true {
		return true
	}
	return false
}

// GetAccountVocab - This function will return the STIX account vocabulary.
func GetAccountVocab() map[string]bool {
	return (map[string]bool{
		"facebook":       true,
		"ldap":           true,
		"nis":            true,
		"openid":         true,
		"radius":         true,
		"skype":          true,
		"tacacs":         true,
		"twitter":        true,
		"unix":           true,
		"windows-local":  true,
		"windows-domain": true,
	})
}

// GetAttackMotivationVocab - This function will return the STIX attack
// motivation vocabulary.
func GetAttackMotivationVocab() map[string]bool {
	return (map[string]bool{
		"accidental":            true,
		"coercion":              true,
		"dominance":             true,
		"ideology":              true,
		"notoriety":             true,
		"organizational-gain":   true,
		"personal-gain":         true,
		"personal-satisfaction": true,
		"revenge":               true,
		"unpredictable":         true,
	})
}

// GetAttackResourceLevelVocab - This function will return the STIX attack
// resource level vocabulary.
func GetAttackResourceLevelVocab() map[string]bool {
	return (map[string]bool{
		"individual":   true,
		"club":         true,
		"contest":      true,
		"team":         true,
		"organization": true,
		"government":   true,
	})
}

// GetEncryptionVocab - This function will return the STIX encryption
// enumeration.
func GetEncryptionVocab() map[string]bool {
	return (map[string]bool{
		"AES-256-GCM":         true,
		"ChaCha20-Poly1305":   true,
		"mime-type-indicated": true,
	})
}

// GetExtensionTypesVocab - This function will return the STIX extension types
// enumeration.
func GetExtensionTypesVocab() map[string]bool {
	return (map[string]bool{
		"new-sdo":                     true,
		"new-sco":                     true,
		"new-sro":                     true,
		"property-extension":          true,
		"toplevel-property-extension": true,
	})
}

// GetGroupingVocab - This function will return the STIX grouping vocabulary
func GetGroupingVocab() map[string]bool {
	return (map[string]bool{
		"suspicious-activity": true,
		"malware-analysis":    true,
		"unspecified":         true,
	})
}

// GetHashingAlgorithmVocab - This function will return the STIX hashing
// algorithm vocabulary
func GetHashingAlgorithmVocab() map[string]bool {
	return (map[string]bool{
		"MD5":      true,
		"SHA-1":    true,
		"SHA-256":  true,
		"SHA-512":  true,
		"SHA3-256": true,
		"SHA3-512": true,
		"SSDEEP":   true,
		"TLSH":     true,
	})
}

// GetIdentityClassVocab - This function will return the STIX identity class
// vocabulary
func GetIdentityClassVocab() map[string]bool {
	return (map[string]bool{
		"individual":   true,
		"group":        true,
		"system":       true,
		"organization": true,
		"class":        true,
		"unknown":      true,
	})
}

// GetImplementationLanguageVocab - This function will return the STIX
// implementation language vocabulary
func GetImplementationLanguageVocab() map[string]bool {
	return (map[string]bool{
		"applescript":  true,
		"bash":         true,
		"c":            true,
		"c++":          true,
		"c#":           true,
		"go":           true,
		"java":         true,
		"javascript":   true,
		"lua":          true,
		"objective-c":  true,
		"perl":         true,
		"php":          true,
		"powershell":   true,
		"python":       true,
		"ruby":         true,
		"scala":        true,
		"swift":        true,
		"typescript":   true,
		"visual-basic": true,
		"x86-32":       true,
		"x86-64":       true,
	})
}

// GetIndicatorTypeVocab - This function will return the STIX indicator type
// vocabulary
func GetIndicatorTypeVocab() map[string]bool {
	return (map[string]bool{
		"anomalous-activity": true,
		"anonymization":      true,
		"benign":             true,
		"compromised":        true,
		"malicious-activity": true,
		"attribution":        true,
		"unknown":            true,
	})
}

// GetIndustrySectorVocab - This function will return the STIX industry sector
// vocabulary
func GetIndustrySectorVocab() map[string]bool {
	return (map[string]bool{
		"agriculture":                   true,
		"aerospace":                     true,
		"automotive":                    true,
		"chemical":                      true,
		"commercial":                    true,
		"communications":                true,
		"construction":                  true,
		"defense":                       true,
		"education":                     true,
		"energy":                        true,
		"entertainment":                 true,
		"financial-services":            true,
		"government":                    true,
		"government-emergency-services": true,
		"government-local":              true,
		"government-national":           true,
		"government-public-services":    true,
		"government-regional":           true,
		"healthcare":                    true,
		"hospitality-leisure":           true,
		"infrastructure":                true,
		"infrastructure-dams":           true,
		"infrastructure-nuclear":        true,
		"infrastructure-water":          true,
		"insurance":                     true,
		"manufacturing":                 true,
		"mining":                        true,
		"non-profit":                    true,
		"pharmaceuticals":               true,
		"retail":                        true,
		"technology":                    true,
		"telecommunications":            true,
		"transportation":                true,
		"utilities":                     true,
	})
}

// GetInfrastructureTypeVocab - This function will return the STIX
// infrastructure type vocabulary
func GetInfrastructureTypeVocab() map[string]bool {
	return (map[string]bool{
		"amplification":        true,
		"anonymization":        true,
		"botnet":               true,
		"command-and-control":  true,
		"control-system":       true,
		"exfiltration":         true,
		"firewall":             true,
		"hosting-malware":      true,
		"hosting-target-lists": true,
		"phishing":             true,
		"reconnaissance":       true,
		"routers-switches":     true,
		"staging":              true,
		"workstation":          true,
		"unknown":              true,
	})
}

// GetMalwareAVResultsVocab - This function will return the STIX malware AV
// results vocabulary
func GetMalwareAVResultsVocab() map[string]bool {
	return (map[string]bool{
		"malicious":  true,
		"suspicious": true,
		"benign":     true,
		"unknown":    true,
	})
}

// GetMalwareCapabilitiesVocab - This function will return the STIX malware
// capabilities vocabulary
func GetMalwareCapabilitiesVocab() map[string]bool {
	return (map[string]bool{
		"accesses-remote-machines":              true,
		"anti-debugging":                        true,
		"anti-disassembly":                      true,
		"anti-emulation":                        true,
		"anti-memory-forensics":                 true,
		"anti-sandbox":                          true,
		"anti-vm":                               true,
		"captures-input-peripherals":            true,
		"captures-output-peripherals":           true,
		"captures-system-state-data":            true,
		"cleans-traces-of-infection":            true,
		"commits-fraud":                         true,
		"communicates-with-c2":                  true,
		"compromises-data-availability":         true,
		"compromises-data-integrity":            true,
		"compromises-system-availability":       true,
		"controls-local-machine":                true,
		"degrades-security-software":            true,
		"degrades-system-updates":               true,
		"determines-c2-server":                  true,
		"emails-spam":                           true,
		"escalates-privileges":                  true,
		"evades-av":                             true,
		"exfiltrates-data":                      true,
		"fingerprints-host":                     true,
		"hides-artifacts":                       true,
		"hides-executing-code":                  true,
		"infects-files":                         true,
		"infects-remote-machines":               true,
		"installs-other-components":             true,
		"persists-after-system-reboot":          true,
		"prevents-artifact-access":              true,
		"prevents-artifact-deletion":            true,
		"probes-network-environment":            true,
		"self-modifies":                         true,
		"steals-authentication-credentials":     true,
		"violates-system-operational-integrity": true,
	})
}

// GetMalwareTypeVocab - This function will return the STIX malware type
// vocabulary
func GetMalwareTypeVocab() map[string]bool {
	return (map[string]bool{
		"adware":                  true,
		"backdoor":                true,
		"bot":                     true,
		"bootkit":                 true,
		"ddos":                    true,
		"downloader":              true,
		"dropper":                 true,
		"exploit-kit":             true,
		"keylogger":               true,
		"ransomware":              true,
		"remote-access-trojan":    true,
		"resource-exploitation":   true,
		"rogue-security-software": true,
		"rootkit":                 true,
		"screen-capture":          true,
		"spyware":                 true,
		"trojan":                  true,
		"unknown":                 true,
		"virus":                   true,
		"webshell":                true,
		"wiper":                   true,
		"worm":                    true,
	})
}

// GetNetworkSocketAddressFamilyVocab - This function will return the STIX
// network socket address family vocabulary
func GetNetworkSocketAddressFamilyVocab() map[string]bool {
	return (map[string]bool{
		"AF_UNSPEC":    true,
		"AF_INET":      true,
		"AF_IPX":       true,
		"AF_APPLETALK": true,
		"AF_NETBIOS":   true,
		"AF_INET6":     true,
		"AF_IRDA":      true,
		"AF_BTH":       true,
	})
}

// GetNetworkSocketTypeVocab - This function will return the STIX network socket
// type family vocabulary
func GetNetworkSocketTypeVocab() map[string]bool {
	return (map[string]bool{
		"SOCK_STREAM":    true,
		"SOC_DGRAM":      true,
		"SOCK_RAW":       true,
		"SOCK_RDM":       true,
		"SOCK_SEQPACKET": true,
	})
}

// GetOpinionVocab - This function will return the STIX opinion vocabulary
func GetOpinionVocab() map[string]bool {
	return (map[string]bool{
		"strongly-disagree": true,
		"disagree":          true,
		"neutral":           true,
		"agree":             true,
		"strongly-agree":    true,
	})
}

// GetPatternTypeVocab - This function will return the STIX pattern type
// vocabulary
func GetPatternTypeVocab() map[string]bool {
	return (map[string]bool{
		"stix":     true,
		"pcre":     true,
		"sigma":    true,
		"snort":    true,
		"suricata": true,
		"yara":     true,
	})
}

// GetProcessorArchitectureVocab - This function will return the STIX processor
// architecture vocabulary
func GetProcessorArchitectureVocab() map[string]bool {
	return (map[string]bool{
		"alpha":   true,
		"arm":     true,
		"ia-64":   true,
		"mips":    true,
		"powerpc": true,
		"sparc":   true,
		"x86":     true,
		"x86-64":  true,
	})
}

// GetRegionVocab - This function will return the STIX region vocabulary
func GetRegionVocab() map[string]bool {
	return (map[string]bool{
		"africa":                  true,
		"eastern-africa":          true,
		"middle-africa":           true,
		"northern-africa":         true,
		"southern-africa":         true,
		"western-africa":          true,
		"americas":                true,
		"caribbean":               true,
		"central-america":         true,
		"latin-america-caribbean": true,
		"northern-america":        true,
		"south-america":           true,
		"asia":                    true,
		"central-asia":            true,
		"eastern-asia":            true,
		"southern-asia":           true,
		"south-eastern-asia":      true,
		"western-asia":            true,
		"europe":                  true,
		"eastern-europe":          true,
		"northern-europe":         true,
		"southern-europe":         true,
		"western-europe":          true,
		"oceania":                 true,
		"antarctica":              true,
		"australia-new-zealand":   true,
		"melanesia":               true,
		"micronesia":              true,
		"polynesia":               true,
	})
}

// GetReportTypeVocab - This function will return the STIX report type
// vocabulary
func GetReportTypeVocab() map[string]bool {
	return (map[string]bool{
		"attack-pattern": true,
		"campaign":       true,
		"identity":       true,
		"indicator":      true,
		"intrusion-set":  true,
		"malware":        true,
		"observed-data":  true,
		"threat-actor":   true,
		"threat-report":  true,
		"tool":           true,
		"vulnerability":  true,
	})
}

// GetThreatActorTypeVocab - This function will return the STIX threat actor
// type vocabulary
func GetThreatActorTypeVocab() map[string]bool {
	return (map[string]bool{
		"activist":            true,
		"competitor":          true,
		"crime-syndicate":     true,
		"criminal":            true,
		"hacker":              true,
		"insider-accidental":  true,
		"insider-disgruntled": true,
		"nation-state":        true,
		"sensationalist":      true,
		"spy":                 true,
		"terrorist":           true,
		"unknown":             true,
	})
}

// GetThreatActorRoleVocab - This function will return the STIX threat actor
// role vocabulary
func GetThreatActorRoleVocab() map[string]bool {
	return (map[string]bool{
		"agent":                    true,
		"director":                 true,
		"independent":              true,
		"infrastructure-architect": true,
		"infrastructure-operator":  true,
		"malware-author":           true,
		"sponsor":                  true,
	})
}

// GetThreatActorSophisticationVocab - This function will return the STIX
// threat actor sophistication vocabulary
func GetThreatActorSophisticationVocab() map[string]bool {
	return (map[string]bool{
		"none":         true,
		"minimal":      true,
		"intermediate": true,
		"advanced":     true,
		"expert":       true,
		"innovator":    true,
		"strategic":    true,
	})
}

// GetToolTypeVocab - This function will return the STIX tool type vocabulary
func GetToolTypeVocab() map[string]bool {
	return (map[string]bool{
		"denial-of-service":       true,
		"exploitation":            true,
		"information-gathering":   true,
		"network-capture":         true,
		"credential-exploitation": true,
		"remote-access":           true,
		"vulnerability-scanning":  true,
		"unknown":                 true,
	})
}

// GetWindowsIntegrityLevelVocab - This function will return the STIX Windows
// integrity level vocabulary
func GetWindowsIntegrityLevelVocab() map[string]bool {
	return (map[string]bool{
		"low":    true,
		"medium": true,
		"high":   true,
		"system": true,
	})
}

// GetWindowsPEBinaryVocab - This function will return the STIX
// Windows PE binary vocabulary
func GetWindowsPEBinaryVocab() map[string]bool {
	return (map[string]bool{
		"dll": true,
		"exe": true,
		"sys": true,
	})
}

// GetWindowsRegistryDatatypeVocab - This function will return the STIX Windows
// registry datatype vocabulary
func GetWindowsRegistryDatatypeVocab() map[string]bool {
	return (map[string]bool{
		"REG_NONE":                       true,
		"REG_SZ":                         true,
		"REG_EXPAND_SZ":                  true,
		"REG_BINARY":                     true,
		"REG_DWORD":                      true,
		"REG_DWORD_BIG_ENDIAN":           true,
		"REG_DWORD_LITTLE_ENDIAN":        true,
		"REG_LINK":                       true,
		"REG_MULTI_SZ":                   true,
		"REG_RESOURCE_LIST":              true,
		"REG_FULL_RESOURCE_DESCRIPTION":  true,
		"REG_RESOURCE_REQUIREMENTS_LIST": true,
		"REG_QWORD":                      true,
		"REG_INVALID_TYPE":               true,
	})
}

// GetWindowsServiceStartTypeVocab - This function will return the STIX
// Windows service start type vocabulary
func GetWindowsServiceStartTypeVocab() map[string]bool {
	return (map[string]bool{
		"SERVICE_AUTO_START":   true,
		"SERVICE_BOOT_START":   true,
		"SERVICE_DEMAND_START": true,
		"SERVICE_DISABLED":     true,
		"SERVICE_SYSTEM_ALERT": true,
	})
}

// GetWindowsServiceTypeVocab - This function will return the STIX
// Windows service type vocabulary
func GetWindowsServiceTypeVocab() map[string]bool {
	return (map[string]bool{
		"SERVICE_KERNEL_DRIVER":       true,
		"SERVICE_FILE_SYSTEM_DRIVER":  true,
		"SERVICE_WIN32_OWN_PROCESS":   true,
		"SERVICE_WIN32_SHARE_PROCESS": true,
	})
}

// GetWindowsServiceStatusVocab - This function will return the STIX
// Windows service status vocabulary
func GetWindowsServiceStatusVocab() map[string]bool {
	return (map[string]bool{
		"SERVICE_CONTINUE_PENDING": true,
		"SERVICE_PAUSE_PENDING":    true,
		"SERVICE_PAUSED":           true,
		"SERVICE_RUNNING":          true,
		"SERVICE_START_PENDING":    true,
		"SERVICE_STOP_PENDING":     true,
		"SERVICE_STOPPED":          true,
	})
}
