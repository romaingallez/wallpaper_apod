package wallpaper

import (
	"fmt"
	"log"

	"github.com/romaingallez/wallpaper_apod/internal/config"
	"github.com/romaingallez/wallpaper_apod/internal/powershell"
)

var (
	PSSetWallPaper = `
	Function Set-WallPaper {

<#

.SYNOPSIS
Applies a specified wallpaper to the current user's desktop

.PARAMETER Image
Provide the exact path to the image

.PARAMETER Style
Provide wallpaper style (Example: Fill, Fit, Stretch, Tile, Center, or Span)

.EXAMPLE
Set-WallPaper -Image "C:\Wallpaper\Default.jpg"
Set-WallPaper -Image "C:\Wallpaper\Background.jpg" -Style Fit

#>

param (
[parameter(Mandatory=$True)]
# Provide path to image
[string]$Image,
# Provide wallpaper style that you would like applied
[parameter(Mandatory=$False)]
[ValidateSet('Fill', 'Fit', 'Stretch', 'Tile', 'Center', 'Span')]
[string]$Style
)

$WallpaperStyle = Switch ($Style) {

"Fill" {"10"}
"Fit" {"6"}
"Stretch" {"2"}
"Tile" {"0"}
"Center" {"0"}
"Span" {"22"}

}

If($Style -eq "Tile") {

New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name WallpaperStyle -PropertyType String -Value $WallpaperStyle -Force
New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name TileWallpaper -PropertyType String -Value 1 -Force

}
Else {

New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name WallpaperStyle -PropertyType String -Value $WallpaperStyle -Force
New-ItemProperty -Path "HKCU:\Control Panel\Desktop" -Name TileWallpaper -PropertyType String -Value 0 -Force

}

Add-Type -TypeDefinition @" 
using System; 
using System.Runtime.InteropServices;

public class Params
{ 
[DllImport("User32.dll",CharSet=CharSet.Unicode)] 
public static extern int SystemParametersInfo (Int32 uAction, 
																						 Int32 uParam, 
																						 String lpvParam, 
																						 Int32 fuWinIni);
}
"@ 

$SPI_SETDESKWALLPAPER = 0x0014
$UpdateIniFile = 0x01
$SendChangeEvent = 0x02

$fWinIni = $UpdateIniFile -bor $SendChangeEvent

$ret = [Params]::SystemParametersInfo($SPI_SETDESKWALLPAPER, 0, $Image, $fWinIni)
}`
)

func SetWallpaper(config config.ConfigType) {
	posh := powershell.NewPS()

	output := fmt.Sprintf("%s\\wallpaper.bmp", config.ImagePath)

	stdOut, stdErr, err := posh.Execute(fmt.Sprintf("%s\n%s", PSSetWallPaper, fmt.Sprintf(`Set-WallPaper -Image "%s" -Style Stretch`, output)))
	if len(stdErr) != 0 {
		log.Println(stdErr)
	}
	if err != nil {
		log.Panic(err)
	}
	log.Println(stdOut)
}
