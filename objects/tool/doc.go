// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package tool implements the STIX 2.1 Tool object.

The following information comes directly from the STIX 2.1 specification.

Tools are legitimate software that can be used by threat actors to perform
attacks. Knowing how and when threat actors use such tools can be important for
understanding how campaigns are executed. Unlike malware, these tools or
software packages are often found on a system and have legitimate purposes for
power users, system administrators, network administrators, or even normal
users. Remote access tools (e.g., RDP) and network scanning tools (e.g., Nmap)
are examples of Tools that may be used by a Threat Actor during an attack.

The Tool SDO characterizes the properties of these software tools and can be
used as a basis for making an assertion about how a Threat Actor uses them
during an attack. It contains properties to name and describe the tool, a list
of Kill Chain Phases the tool can be used to carry out, and the version of the
tool.

This SDO MUST NOT be used to characterize malware. Further, Tool MUST NOT be
used to characterize tools used as part of a course of action in response to an
attack.
*/
package tool
