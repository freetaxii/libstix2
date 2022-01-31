// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package vocabs implements the STIX 2 Vocabularies.

This package defines functions for working with STIX vocabularies.

The following sections provide object-specific listings for each of the
vocabularies referenced in the object description sections defined in Sections
4, 5, 6, and 7.

STIX vocabularies that have type names ending in '-ov', are "open": they provide
a listing of common and industry accepted terms as a guide to the user but do
not limit the user to that defined list. These vocabularies are referenced from
the STIX Objects as type open-vocab and have a statement indicating which
vocabulary should be used.

STIX vocabularies that have type names ending in '-enum' are "closed": the only
valid values are those in the vocabulary. These vocabularies are referenced
from the STIX Objects as type enum and have a statement indicating which
enumeration must be used.




Account Type Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: account-type-ov

The account type vocabulary is currently used in the following SCOs:
 * User Account

An open vocabulary of User Account types.




Attack Motivation Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: attack-motivation-ov

The attack motivation vocabulary is currently used in the following SDOs:
 * Intrusion Set
 * Threat Actor

Knowing a Threat Actor or Intrusion Set's motivation may allow an analyst or
defender to better understand likely targets and behaviors.

Motivation shapes the intensity and the persistence of an attack. Threat Actors
and Intrusion Sets usually act in a manner that reflects their underlying
emotion or situation, and this informs defenders of the manner of attack. For
example, a spy motivated by nationalism (ideology) likely has the patience to
achieve long-term goals and work quietly for years, whereas a cyber-vandal out
for notoriety can create an intense and attention-grabbing attack but may
quickly lose interest and move on. Understanding these differences allows
defenders to implement controls tailored to each type of attack for greatest
efficiency.

This section including vocabulary items and their descriptions is based on the
Threat Agent Motivations publication from Intel Corp in February 2015
[Casey 2015].




Attack Resource Level Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: attack-resource-level-ov

The attack resource level vocabulary is currently used in the following SDO(s):
 * Intrusion Set
 * Threat Actor

Attack Resource Level is an open vocabulary that captures the general level of
resources that a threat actor, intrusion set, or campaign might have access to.
It ranges from individual, a person acting alone, to government, the resources
of a national government.

This section including vocabulary items and their descriptions is based on the
Threat Agent Library publication from Intel Corp in September 2007
[Casey 2007].




Encryption Algorithm Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: encryption-algorithm-enum

The encryption algorithm enumeration is currently used in the following SCOs:
 * Artifact

An enumeration of encryption algorithms for sharing defanged and/or confidential
artifacts.




Extension Types Enumeration

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: extension-type-enum

The Extensions Type enumeration is used in the Extension meta-object.




Grouping Context Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: grouping-context-ov

The Grouping Context open vocabulary is currently used in the following object:
 * Grouping

While the majority of this vocabulary is undefined (producers may use custom
vocabulary entries), it has been added specifically to capture the
suspicious-activity-event value. That value indicates that the information
contained in the Grouping relates to a suspicious event.




Hashing Algorithm Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: hash-algorithm-ov

The Hashing Algorithm open vocabulary is currently used in the following
objects:
 * External Reference
 * Artifact
 * File
 * Alternate Data Stream
 * Windows™ PE Binary File
 * Windows™ PE Optional Header
 * Windows™ PE Section
 * X.509 Certificate

A vocabulary of hashing algorithms.




Identity Class Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: identity-class-ov

The identity class vocabulary is currently used in the following SDO(s):
 * Identity

This vocabulary describes the type of entity that the Identity represents:
whether it describes an organization, group, individual, or class.




Implementation Language Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: implementation-language-ov

The implementation language vocabulary is currently used in the following SDO(s):
 * Malware

This is a non-exhaustive, open vocabulary that covers common programming
languages and is intended to characterize the languages that may have been used
to implement a malware instance or family.




Indicator Type Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: indicator-type-ov

The indicator type vocabulary is currently used in the following SDO(s):
 * Indicator

Indicator type is an open vocabulary used to categorize Indicators. It is
intended to be high-level to promote consistent practices. Indicator types
should not be used to capture information that can be better captured via
related Malware or Attack Pattern objects. It is better to link an Indicator to
a Malware object describing Poison Ivy rather than simply providing a type or
label of "poison-ivy".




Industry Sector Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: industry-sector-ov

The industry sector vocabulary is currently used in the following SDO(s):
 * Identity

Industry sector is an open vocabulary that describes industrial and commercial
sectors. It is intended to be holistic; it has been derived from several other
lists and is not limited to "critical infrastructure" sectors.




Infrastructure Type Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: infrastructure-type-ov

The infrastructure type vocabulary is currently used in the following SDO(s):
 * Infrastructure

A non-exhaustive enumeration of infrastructure types.




Malware Result Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: malware-result-ov

The processor architecture vocabulary is currently used in the following SDO(s):
 * Malware Analysis

This is a non-exhaustive, open vocabulary that captures common types of scanner
or tool analysis process results.





Malware Capabilities Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: malware-capabilities-ov

The malware capabilities vocabulary is currently used in the following SDO(s):
 * Malware

This is an open vocabulary that covers common capabilities that may be exhibited
by a malware instance or family.




Malware Type Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: malware-type-ov

The malware type vocabulary is currently used in the following SDO(s):
 * Malware

Malware type is an open vocabulary that represents different types and functions
of malware. Malware types are not mutually exclusive; for example, a malware
instance can be both spyware and a screen capture tool.




Network Socket Address Family Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: network-socket-address-family-enum

The network socket address family vocabulary is currently used in the following SCO(s):
 * Network Traffic (Network Socket extension)

An enumeration of network socket address family types.




Network Socket Type Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: network-socket-type-enum

The network socket type vocabulary is currently used in the following SCO(s):
 * Network Traffic (Network Socket extension)

An enumeration of network socket types.




Opinion Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: opinion-enum

The agreement enumeration is currently used in the following SDOs:
 * Opinion

This enumeration captures a degree of agreement with the information in a STIX
Object. It is an ordered enumeration, with the earlier terms representing
disagreement, the middle term neutral, and the later terms representing
agreement.




Pattern Type Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: pattern-type-ov

The pattern type vocabulary is currently used in the following SDO(s):
 * Indicator

This is a non-exhaustive, open vocabulary that covers common pattern languages
and is intended to characterize the pattern language that the indicator pattern
is expressed in.




Processor Architecture Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: processor-architecture-ov

The processor architecture vocabulary is currently used in the following SDO(s):
 * Malware

This is a non-exhaustive, open vocabulary that covers common processor
architectures and is intended to characterize the architectures that a malware
instance or family may be able to execute on.




Region Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: region-ov

The region vocabulary is currently used in the following SDO(s):
 * Location

A list of world regions based on the United Nations geoscheme [UNSD M49].




Report Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: report-type-ov

The report type vocabulary is currently used in the following SDO(s):
 * Report

Report type is an open vocabulary to describe the primary purpose or subject of
a report. For example, a report that contains malware and indicators for that
malware should have a report type of malware to capture that the malware is the
primary purpose. Report types are not mutually exclusive: a Report can be both
a malware report and a tool report. Just because a report contains objects of a
type does not mean that the report should include that type. If the objects are
there to simply provide evidence or context for other objects, it is not
necessary to include them in the type.




Threat Actor Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: threat-actor-type-ov

The threat actor type vocabulary is currently used in the following SDO(s):
 * Threat Actor

Threat actor type is an open vocabulary used to describe what type of threat
actor the individual or group is. For example, some threat actors are
competitors who try to steal information, while others are activists who act in
support of a social or political cause. Actor types are not mutually exclusive:
a threat actor can be both a disgruntled insider and a spy. [Casey 2007])




Threat Actor Role Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: threat-actor-role-ov

The threat actor role vocabulary is currently used in the following SDO(s):
 * Threat Actor

Threat actor role is an open vocabulary that is used to describe the different
roles that a threat actor can play. For example, some threat actors author
malware or operate botnets while other actors actually carry out attacks
directly.

Threat actor roles are not mutually exclusive. For example, an actor can be both
a financial backer for attacks and also direct attacks.




Threat Actor Sophistication Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: threat-actor-sophistication-ov

Threat actor sophistication vocabulary is currently used in the following SDO(s):
 * Threat Actor

Threat actor sophistication vocabulary captures the skill level of a threat
actor. It ranges from "none", which describes a complete novice,
to "strategic", which describes an attacker who is able to influence supply
chains to introduce vulnerabilities. This vocabulary is separate from resource
level because an innovative, highly-skilled threat actor may have access to
very few resources while a minimal-level actor might have the resources of an
organized crime ring.




Tool Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: tool-type-ov

The tool type vocabulary is currently used in the following SDO(s):
 * Tool

Tool types describe the categories of tools that can be used to perform
attacks.




Windows™ Integrity Level Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: windows-integrity-level-enum

The Windows integrity level enumeration is currently used in the following STIX
Cyber-observable Object(s):
 * Process (Windows Process extension)

Windows integrity levels are a security feature and represent the
trustworthiness of an object.




Windows™ PE Binary Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: windows-pebinary-type-ov

The Windows PE binary vocabulary is currently used in the following SCO(s):
 * File (Windows PE Binary extension)

An open vocabulary of Windows PE binary types.




Windows™ Registry Datatype Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: windows-registry-datatype-enum

The Windows registry datatype vocabulary is currently used in the following SCO(s):
 * Windows Registry Key

An enumeration of Windows registry data types.




Windows™ Service Start Type Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: windows-service-start-type-enum

The Windows service start type vocabulary is currently used in the following SCO(s):
 * Process (Windows Service extension)

An enumeration of Windows service start types.




Windows™ Service Type Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: windows-service-type-enum

The Windows service type vocabulary is currently used in the following SCO(s):
 * Process (Windows Service extension)

An enumeration of Windows service types.




Windows™ Service Status Enumeration

The following information comes directly from the STIX 2 specification documents.

Enumeration Name: windows-service-status-enum

The Windows service status vocabulary is currently used in the following SCO(s):
 * Process (Windows Service extension)

An enumeration of Windows service statuses.
*/
package vocabs
