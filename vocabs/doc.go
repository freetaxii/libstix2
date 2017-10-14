// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

/*
Package vocabs implements the STIX 2 Vocabularies.
This package defines variables that contain all of the values for each vocabulary.

The following information comes directly from the STIX 2 specification documents.

The following sections provide object-specific listings for each of the
vocabularies referenced in the object description sections defined in STIX™
Version 2.0. Part 2: STIX Objects. STIX vocabularies, which all have type names
ending in '-ov', are "open": they provide a listing of common and industry
accepted terms as a guide to the user but do not limit the user to that defined
list.




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
Threat Agent Library publication from Intel Corp in September 2007 [Casey 2007].




Identity Class Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: identity-class-ov

The identity class vocabulary is currently used in the following SDO(s):
 * Identity

This vocabulary describes the type of entity that the Identity represents:
whether it describes an organization, group, individual, or class.




Indicator Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: indicator-label-ov

The indicator label vocabulary is currently used in the following SDO(s):
 * Indicator

Indicator labels is an open vocabulary used to categorize Indicators. It is
intended to be high-level to promote consistent practices. Indicator labels
should not be used to capture information that can be better captured via
related Malware or Attack Pattern objects. It is better to link an Indicator to
a Malware object describing Poison Ivy rather than simply labeling it with
"poison-ivy".




Industry Sector Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: industry-sector-ov

The industry sector vocabulary is currently used in the following SDO(s):
 * Identity

Industry sector is an open vocabulary that describes industrial and commercial
sectors. It is intended to be holistic; it has been derived from several other
lists and is not limited to "critical infrastructure" sectors.




Malware Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: malware-label-ov

The malware label vocabulary is currently used in the following SDO(s):
 * Malware

Malware label is an open vocabulary that represents different types and
functions of malware. Malware labels are not mutually exclusive; a malware
instance can be both spyware and a screen capture tool.




Report Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: report-label-ov

The report label vocabulary is currently used in the following SDO(s):
 * Report

Report label is an open vocabulary to describe the primary purpose or subject of
a report. For example, a report that contains malware and indicators for that
malware should have a report label of malware to capture that the malware is the
primary purpose. Report labels are not mutually exclusive: a Report can be both
a malware report and a tool report. Just because a report contains objects of a
type does not mean that the report should include that label.  If the objects
are there to simply provide evidence or context for other objects, it is not
necessary to include them in the label.




Threat Actor Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: threat-actor-label-ov

The threat actor label vocabulary is currently used in the following SDO(s):
 * Threat Actor

Threat actor label is an open vocabulary used to describe what type of threat
actor the individual or group is. For example, some threat actors are
competitors who try to steal information, while others are activists who act in
support of a social or political cause. Actor labels are not mutually exclusive:
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
actor. It ranges from "none", which describes a complete novice, to "strategic",
which describes an attacker who is able to influence supply chains to introduce
vulnerabilities. This vocabulary is separate from resource level because an
innovative, highly-skilled threat actor may have access to very few resources
while a minimal-level actor might have the resources of an organized crime ring.




Tool Label Vocabulary

The following information comes directly from the STIX 2 specification documents.

Vocabulary Name: tool-label-ov

The tool label vocabulary is currently used in the following SDO(s):
 * Tool

Tool labels describe the categories of tools that can be used to perform attacks.
*/
package vocabs
