package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/automate-powerautomate/magicapp"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: automate-powerautomate
description: Describe the main purpose of this kitchen
---

# automate-powerautomate

Hello, this is Bing. I can help you with some information and resources on how to use PowerShell for managing Power Automate. 😊

PowerShell is a scripting language that allows you to automate tasks and interact with various Microsoft services, including Power Platform. With PowerShell cmdlets for Power Platform administrators, you can manage and administer Power Automate flows, Power Apps, and environments¹.

To get started with PowerShell for Power Platform, you need to install the necessary modules on your local machine and sign in to your Power Platform account using the Add-PowerAppsAccount cmdlet². You can also use a client ID and secret or certificate to sign in non-interactively².

Once you are signed in, you can use the PowerShell cmdlets to perform various operations on your Power Automate resources, such as creating, updating, deleting, starting, and stopping flows. You can also list, export, and import flow definitions, as well as get flow run history and status³.

For more details and examples on how to use PowerShell for Power Automate, you can check out these links:

- [PowerShell support for Power Apps and Power Automate](^1^)
- [Get started with PowerShell for Power Platform Administrators](^2^)
- [Execute PowerShell scripts through Power Automate flows](^3^)

I hope this helps you learn more about PowerShell and Power Automate. Have a great day! 😊

Source: Conversation with Bing, 26/01/2024
(1) PowerShell support for Power Apps and Power Automate. https://learn.microsoft.com/en-us/power-platform/admin/powerapps-powershell.
(2) Get started with PowerShell for Power Platform Administrators. https://learn.microsoft.com/en-us/power-platform/admin/powershell-getting-started.
(3) Execute Power Shell scripts through Power Automate flows. https://powerusers.microsoft.com/t5/Building-Flows/Execute-Power-Shell-scripts-through-Power-Automate-flows/td-p/731472.
(4) PowerShell support for Power Apps and Power Automate. https://learn.microsoft.com/en-us/power-platform/admin/powerapps-powershell.
(5) Get started with PowerShell for Power Platform Administrators. https://learn.microsoft.com/en-us/power-platform/admin/powershell-getting-started.
(6) Execute Power Shell scripts through Power Automate flows. https://powerusers.microsoft.com/t5/Building-Flows/Execute-Power-Shell-scripts-through-Power-Automate-flows/td-p/731472.`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("automate-powerautomate", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.Execute(name, "automate-powerautomate", "")
}
