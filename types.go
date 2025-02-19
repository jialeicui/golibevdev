package golibevdev

// Input event codes from linux/input-event-codes.h

//go:generate stringer -type=EventType
type EventType uint16

const (
	EvSyn EventType = iota
	EvKey
	EvRel
	EvAbs
	EvMsc
	EvSw
	EvLed      = 0x11
	EvSnd      = 0x12
	EvRep      = 0x14
	EvFf       = 0x15
	EvPwr      = 0x16
	EvFfStatus = 0x17
	EvMax      = 0x1f
	EvCnt      = EvMax + 1
)

type EventCode interface {
	Type() string
	Value() uint16
}

//go:generate stringer -type=KeyEventCode
type KeyEventCode uint16

func (c KeyEventCode) Type() string {
	return "Key"
}

func (c KeyEventCode) Value() uint16 {
	return uint16(c)
}

const (
	KeyReserved         KeyEventCode = 0
	KeyEsc              KeyEventCode = 1
	Key1                KeyEventCode = 2
	Key2                KeyEventCode = 3
	Key3                KeyEventCode = 4
	Key4                KeyEventCode = 5
	Key5                KeyEventCode = 6
	Key6                KeyEventCode = 7
	Key7                KeyEventCode = 8
	Key8                KeyEventCode = 9
	Key9                KeyEventCode = 10
	Key0                KeyEventCode = 11
	KeyMinus            KeyEventCode = 12
	KeyEqual            KeyEventCode = 13
	KeyBackspace        KeyEventCode = 14
	KeyTab              KeyEventCode = 15
	KeyQ                KeyEventCode = 16
	KeyW                KeyEventCode = 17
	KeyE                KeyEventCode = 18
	KeyR                KeyEventCode = 19
	KeyT                KeyEventCode = 20
	KeyY                KeyEventCode = 21
	KeyU                KeyEventCode = 22
	KeyI                KeyEventCode = 23
	KeyO                KeyEventCode = 24
	KeyP                KeyEventCode = 25
	KeyLeftBrace        KeyEventCode = 26
	KeyRightBrace       KeyEventCode = 27
	KeyEnter            KeyEventCode = 28
	KeyLeftCtrl         KeyEventCode = 29
	KeyA                KeyEventCode = 30
	KeyS                KeyEventCode = 31
	KeyD                KeyEventCode = 32
	KeyF                KeyEventCode = 33
	KeyG                KeyEventCode = 34
	KeyH                KeyEventCode = 35
	KeyJ                KeyEventCode = 36
	KeyK                KeyEventCode = 37
	KeyL                KeyEventCode = 38
	KeySemicolon        KeyEventCode = 39
	KeyApostrophe       KeyEventCode = 40
	KeyGrave            KeyEventCode = 41
	KeyLeftShift        KeyEventCode = 42
	KeyBackslash        KeyEventCode = 43
	KeyZ                KeyEventCode = 44
	KeyX                KeyEventCode = 45
	KeyC                KeyEventCode = 46
	KeyV                KeyEventCode = 47
	KeyB                KeyEventCode = 48
	KeyN                KeyEventCode = 49
	KeyM                KeyEventCode = 50
	KeyComma            KeyEventCode = 51
	KeyDot              KeyEventCode = 52
	KeySlash            KeyEventCode = 53
	KeyRightShift       KeyEventCode = 54
	KeyKPAsterisk       KeyEventCode = 55
	KeyLeftAlt          KeyEventCode = 56
	KeySpace            KeyEventCode = 57
	KeyCapsLock         KeyEventCode = 58
	KeyF1               KeyEventCode = 59
	KeyF2               KeyEventCode = 60
	KeyF3               KeyEventCode = 61
	KeyF4               KeyEventCode = 62
	KeyF5               KeyEventCode = 63
	KeyF6               KeyEventCode = 64
	KeyF7               KeyEventCode = 65
	KeyF8               KeyEventCode = 66
	KeyF9               KeyEventCode = 67
	KeyF10              KeyEventCode = 68
	KeyNumLock          KeyEventCode = 69
	KeyScrollLock       KeyEventCode = 70
	KeyKP7              KeyEventCode = 71
	KeyKP8              KeyEventCode = 72
	KeyKP9              KeyEventCode = 73
	KeyKPMinus          KeyEventCode = 74
	KeyKP4              KeyEventCode = 75
	KeyKP5              KeyEventCode = 76
	KeyKP6              KeyEventCode = 77
	KeyKPPlus           KeyEventCode = 78
	KeyKP1              KeyEventCode = 79
	KeyKP2              KeyEventCode = 80
	KeyKP3              KeyEventCode = 81
	KeyKP0              KeyEventCode = 82
	KeyKPDot            KeyEventCode = 83
	KeyZenkakuHankaku   KeyEventCode = 85
	Key102nd            KeyEventCode = 86
	KeyF11              KeyEventCode = 87
	KeyF12              KeyEventCode = 88
	KeyRo               KeyEventCode = 89
	KeyKatakana         KeyEventCode = 90
	KeyHiragana         KeyEventCode = 91
	KeyHenkan           KeyEventCode = 92
	KeyKatakanaHiragana KeyEventCode = 93
	KeyMuhenkan         KeyEventCode = 94
	KeyKPJpComma        KeyEventCode = 95
	KeyKPEnter          KeyEventCode = 96
	KeyRightCtrl        KeyEventCode = 97
	KeyKPSlash          KeyEventCode = 98
	KeySysRq            KeyEventCode = 99
	KeyRightAlt         KeyEventCode = 100
	KeyLineFeed         KeyEventCode = 101
	KeyHome             KeyEventCode = 102
	KeyUp               KeyEventCode = 103
	KeyPageUp           KeyEventCode = 104
	KeyLeft             KeyEventCode = 105
	KeyRight            KeyEventCode = 106
	KeyEnd              KeyEventCode = 107
	KeyDown             KeyEventCode = 108
	KeyPageDown         KeyEventCode = 109
	KeyInsert           KeyEventCode = 110
	KeyDelete           KeyEventCode = 111
	KeyMacro            KeyEventCode = 112
	KeyMute             KeyEventCode = 113
	KeyVolumeDown       KeyEventCode = 114
	KeyVolumeUp         KeyEventCode = 115
	KeyPower            KeyEventCode = 116
	KeyKPEqual          KeyEventCode = 117
	KeyKPPlusMinus      KeyEventCode = 118
	KeyPause            KeyEventCode = 119
	KeyScale            KeyEventCode = 120
	KeyKPComma          KeyEventCode = 121
	KeyHangeul          KeyEventCode = 122
	KeyHanguel          KeyEventCode = KeyHangeul
	KeyHanja            KeyEventCode = 123
	KeyYen              KeyEventCode = 124
	KeyLeftMeta         KeyEventCode = 125
	KeyRightMeta        KeyEventCode = 126
	KeyCompose          KeyEventCode = 127
	KeyStop             KeyEventCode = 128
	KeyAgain            KeyEventCode = 129
	KeyProps            KeyEventCode = 130
	KeyUndo             KeyEventCode = 131
	KeyFront            KeyEventCode = 132
	KeyCopy             KeyEventCode = 133
	KeyOpen             KeyEventCode = 134
	KeyPaste            KeyEventCode = 135
	KeyFind             KeyEventCode = 136
	KeyCut              KeyEventCode = 137
	KeyHelp             KeyEventCode = 138
	KeyMenu             KeyEventCode = 139
	KeyCalc             KeyEventCode = 140
	KeySetup            KeyEventCode = 141
	KeySleep            KeyEventCode = 142
	KeyWakeUp           KeyEventCode = 143
	KeyFile             KeyEventCode = 144
	KeySendFile         KeyEventCode = 145
	KeyDeleteFile       KeyEventCode = 146
	KeyXfer             KeyEventCode = 147
	KeyProg1            KeyEventCode = 148
	KeyProg2            KeyEventCode = 149
	KeyWWW              KeyEventCode = 150
	KeyMSDOS            KeyEventCode = 151
	KeyCoffee           KeyEventCode = 152
	KeyScreenLock                    = KeyCoffee
	KeyRotateDisplay    KeyEventCode = 153
	KeyDirection                     = KeyRotateDisplay
	KeyCycleWindows     KeyEventCode = 154
	KeyMail             KeyEventCode = 155
	KeyBookmarks        KeyEventCode = 156
	KeyComputer         KeyEventCode = 157
	KeyBack             KeyEventCode = 158
	KeyForward          KeyEventCode = 159
	KeyCloseCD          KeyEventCode = 160
	KeyEjectCD          KeyEventCode = 161
	KeyEjectCloseCD     KeyEventCode = 162
	KeyNextSong         KeyEventCode = 163
	KeyPlayPause        KeyEventCode = 164
	KeyPreviousSong     KeyEventCode = 165
	KeyStopCD           KeyEventCode = 166
	KeyRecord           KeyEventCode = 167
	KeyRewind           KeyEventCode = 168
	KeyPhone            KeyEventCode = 169
	KeyISO              KeyEventCode = 170
	KeyConfig           KeyEventCode = 171
	KeyHomePage         KeyEventCode = 172
	KeyRefresh          KeyEventCode = 173
	KeyExit             KeyEventCode = 174
	KeyMove             KeyEventCode = 175
	KeyEdit             KeyEventCode = 176
	KeyScrollUp         KeyEventCode = 177
	KeyScrollDown       KeyEventCode = 178
	KeyKPLeftParen      KeyEventCode = 179
	KeyKPRightParen     KeyEventCode = 180
	KeyNew              KeyEventCode = 181
	KeyRedo             KeyEventCode = 182
	KeyF13              KeyEventCode = 183
	KeyF14              KeyEventCode = 184
	KeyF15              KeyEventCode = 185
	KeyF16              KeyEventCode = 186
	KeyF17              KeyEventCode = 187
	KeyF18              KeyEventCode = 188
	KeyF19              KeyEventCode = 189
	KeyF20              KeyEventCode = 190
	KeyF21              KeyEventCode = 191
	KeyF22              KeyEventCode = 192
	KeyF23              KeyEventCode = 193
	KeyF24              KeyEventCode = 194
	KeyPlayCD           KeyEventCode = 200
	KeyPauseCD          KeyEventCode = 201
	KeyProg3            KeyEventCode = 202
	KeyProg4            KeyEventCode = 203
	KeyAllApps          KeyEventCode = 204
	KeyDashboard        KeyEventCode = KeyAllApps
	KeySuspend          KeyEventCode = 205
	KeyClose            KeyEventCode = 206
	KeyPlay             KeyEventCode = 207
	KeyFastForward      KeyEventCode = 208
	KeyBassBoost        KeyEventCode = 209
	KeyPrint            KeyEventCode = 210
	KeyHP               KeyEventCode = 211
	KeyCamera           KeyEventCode = 212
	KeySound            KeyEventCode = 213
	KeyQuestion         KeyEventCode = 214
	KeyEmail            KeyEventCode = 215
	KeyChat             KeyEventCode = 216
	KeySearch           KeyEventCode = 217
	KeyConnect          KeyEventCode = 218
	KeyFinance          KeyEventCode = 219
	KeySport            KeyEventCode = 220
	KeyShop             KeyEventCode = 221
	KeyAlterase         KeyEventCode = 222
	KeyCancel           KeyEventCode = 223
	KeyBrightnessDown   KeyEventCode = 224
	KeyBrightnessUp     KeyEventCode = 225
	KeyMedia            KeyEventCode = 226
	KeySwitchVideoMode  KeyEventCode = 227
	KeyKbdIllumToggle   KeyEventCode = 228
	KeyKbdIllumDown     KeyEventCode = 229
	KeyKbdIllumUp       KeyEventCode = 230
	KeySend             KeyEventCode = 231
	KeyReply            KeyEventCode = 232
	KeyForwardMail      KeyEventCode = 233
	KeySave             KeyEventCode = 234
	KeyDocuments        KeyEventCode = 235
	KeyBattery          KeyEventCode = 236
	KeyBluetooth        KeyEventCode = 237
	KeyWLAN             KeyEventCode = 238
	KeyUWB              KeyEventCode = 239
	KeyUnknown          KeyEventCode = 240
	KeyVideoNext        KeyEventCode = 241
	KeyVideoPrev        KeyEventCode = 242
	KeyBrightnessCycle  KeyEventCode = 243
	KeyBrightnessAuto   KeyEventCode = 244
	KeyBrightnessZero   KeyEventCode = KeyBrightnessAuto
	KeyDisplayOff       KeyEventCode = 245
	KeyWWAN             KeyEventCode = 246
	KeyWiMax            KeyEventCode = KeyWWAN
	KeyRFKill           KeyEventCode = 247
	KeyMicMute          KeyEventCode = 248

	BtnMisc KeyEventCode = 0x100
	Btn0    KeyEventCode = 0x100
	Btn1    KeyEventCode = 0x101
	Btn2    KeyEventCode = 0x102
	Btn3    KeyEventCode = 0x103
	Btn4    KeyEventCode = 0x104
	Btn5    KeyEventCode = 0x105
	Btn6    KeyEventCode = 0x106
	Btn7    KeyEventCode = 0x107
	Btn8    KeyEventCode = 0x108
	Btn9    KeyEventCode = 0x109

	BtnMouse   KeyEventCode = 0x110
	BtnLeft    KeyEventCode = 0x110
	BtnRight   KeyEventCode = 0x111
	BtnMiddle  KeyEventCode = 0x112
	BtnSide    KeyEventCode = 0x113
	BtnExtra   KeyEventCode = 0x114
	BtnForward KeyEventCode = 0x115
	BtnBack    KeyEventCode = 0x116
	BtnTask    KeyEventCode = 0x117

	BtnJoystick KeyEventCode = 0x120
	BtnTrigger  KeyEventCode = 0x120
	BtnThumb    KeyEventCode = 0x121
	BtnThumb2   KeyEventCode = 0x122
	BtnTop      KeyEventCode = 0x123
	BtnTop2     KeyEventCode = 0x124
	BtnPinkie   KeyEventCode = 0x125
	BtnBase     KeyEventCode = 0x126
	BtnBase2    KeyEventCode = 0x127
	BtnBase3    KeyEventCode = 0x128
	BtnBase4    KeyEventCode = 0x129
	BtnBase5    KeyEventCode = 0x12a
	BtnBase6    KeyEventCode = 0x12b
	BtnDead     KeyEventCode = 0x12f

	BtnGamepad KeyEventCode = 0x130
	BtnSouth   KeyEventCode = 0x130
	BtnA       KeyEventCode = BtnSouth
	BtnEast    KeyEventCode = 0x131
	BtnB       KeyEventCode = BtnEast
	BtnC       KeyEventCode = 0x132
	BtnNorth   KeyEventCode = 0x133
	BtnX       KeyEventCode = BtnNorth
	BtnWest    KeyEventCode = 0x134
	BtnY       KeyEventCode = BtnWest
	BtnZ       KeyEventCode = 0x135
	BtnTL      KeyEventCode = 0x136
	BtnTR      KeyEventCode = 0x137
	BtnTL2     KeyEventCode = 0x138
	BtnTR2     KeyEventCode = 0x139
	BtnSelect  KeyEventCode = 0x13a
	BtnStart   KeyEventCode = 0x13b
	BtnMode    KeyEventCode = 0x13c
	BtnThumbL  KeyEventCode = 0x13d
	BtnThumbR  KeyEventCode = 0x13e

	BtnDigi          KeyEventCode = 0x140
	BtnToolPen       KeyEventCode = 0x140
	BtnToolRubber    KeyEventCode = 0x141
	BtnToolBrush     KeyEventCode = 0x142
	BtnToolPencil    KeyEventCode = 0x143
	BtnToolAirbrush  KeyEventCode = 0x144
	BtnToolFinger    KeyEventCode = 0x145
	BtnToolMouse     KeyEventCode = 0x146
	BtnToolLens      KeyEventCode = 0x147
	BtnToolQuintTap  KeyEventCode = 0x148
	BtnStylus3       KeyEventCode = 0x149
	BtnTouch         KeyEventCode = 0x14a
	BtnStylus        KeyEventCode = 0x14b
	BtnStylus2       KeyEventCode = 0x14c
	BtnToolDoubleTap KeyEventCode = 0x14d
	BtnToolTripleTap KeyEventCode = 0x14e
	BtnToolQuadTap   KeyEventCode = 0x14f

	BtnWheel    KeyEventCode = 0x150
	BtnGearDown KeyEventCode = 0x150
	BtnGearUp   KeyEventCode = 0x151

	KeyOk                      KeyEventCode = 0x160
	KeySelect                  KeyEventCode = 0x161
	KeyGoto                    KeyEventCode = 0x162
	KeyClear                   KeyEventCode = 0x163
	KeyPower2                  KeyEventCode = 0x164
	KeyOption                  KeyEventCode = 0x165
	KeyInfo                    KeyEventCode = 0x166
	KeyTime                    KeyEventCode = 0x167
	KeyVendor                  KeyEventCode = 0x168
	KeyArchive                 KeyEventCode = 0x169
	KeyProgram                 KeyEventCode = 0x16a
	KeyChannel                 KeyEventCode = 0x16b
	KeyFavorites               KeyEventCode = 0x16c
	KeyEPG                     KeyEventCode = 0x16d
	KeyPVR                     KeyEventCode = 0x16e
	KeyMHP                     KeyEventCode = 0x16f
	KeyLanguage                KeyEventCode = 0x170
	KeyTitle                   KeyEventCode = 0x171
	KeySubtitle                KeyEventCode = 0x172
	KeyAngle                   KeyEventCode = 0x173
	KeyFullScreen              KeyEventCode = 0x174
	KeyZoom                    KeyEventCode = KeyFullScreen
	KeyMode                    KeyEventCode = 0x175
	KeyKeyboard                KeyEventCode = 0x176
	KeyAspectRatio             KeyEventCode = 0x177
	KeyScreen                  KeyEventCode = KeyAspectRatio
	KeyPC                      KeyEventCode = 0x178
	KeyTV                      KeyEventCode = 0x179
	KeyTV2                     KeyEventCode = 0x17a
	KeyVCR                     KeyEventCode = 0x17b
	KeyVCR2                    KeyEventCode = 0x17c
	KeyCD                      KeyEventCode = 0x17f
	KeyTape                    KeyEventCode = 0x180
	KeyRadio                   KeyEventCode = 0x181
	KeyTuner                   KeyEventCode = 0x182
	KeyPlayer                  KeyEventCode = 0x183
	KeyText                    KeyEventCode = 0x184
	KeyDVD                     KeyEventCode = 0x185
	KeyAUX                     KeyEventCode = 0x186
	KeyMP3                     KeyEventCode = 0x187
	KeyAudio                   KeyEventCode = 0x188
	KeyVideo                   KeyEventCode = 0x189
	KeyDirectory               KeyEventCode = 0x18a
	KeyList                    KeyEventCode = 0x18b
	KeyMemo                    KeyEventCode = 0x18c
	KeyCalendar                KeyEventCode = 0x18d
	KeyRed                     KeyEventCode = 0x18e
	KeyGreen                   KeyEventCode = 0x18f
	KeyYellow                  KeyEventCode = 0x190
	KeyBlue                    KeyEventCode = 0x191
	KeyChannelUp               KeyEventCode = 0x192
	KeyChannelDown             KeyEventCode = 0x193
	KeyFirst                   KeyEventCode = 0x194
	KeyLast                    KeyEventCode = 0x195
	KeyAB                      KeyEventCode = 0x196
	KeyNext                    KeyEventCode = 0x197
	KeyRestart                 KeyEventCode = 0x198
	KeySlow                    KeyEventCode = 0x199
	KeyShuffle                 KeyEventCode = 0x19a
	KeyBreak                   KeyEventCode = 0x19b
	KeyPrevious                KeyEventCode = 0x19c
	KeyDigits                  KeyEventCode = 0x19d
	KeyTeen                    KeyEventCode = 0x19e
	KeyTwen                    KeyEventCode = 0x19f
	KeyVideoPhone              KeyEventCode = 0x1a0
	KeyGames                   KeyEventCode = 0x1a1
	KeyZoomIn                  KeyEventCode = 0x1a2
	KeyZoomOut                 KeyEventCode = 0x1a3
	KeyZoomReset               KeyEventCode = 0x1a4
	KeyWordProcessor           KeyEventCode = 0x1a5
	KeyEditor                  KeyEventCode = 0x1a6
	KeySpreadsheet             KeyEventCode = 0x1a7
	KeyGraphicsEditor          KeyEventCode = 0x1a8
	KeyPresentation            KeyEventCode = 0x1a9
	KeyDatabase                KeyEventCode = 0x1aa
	KeyNews                    KeyEventCode = 0x1ab
	KeyVoicemail               KeyEventCode = 0x1ac
	KeyAddressBook             KeyEventCode = 0x1ad
	KeyMessenger               KeyEventCode = 0x1ae
	KeyDisplayToggle           KeyEventCode = 0x1af
	KeyBrightnessToggle        KeyEventCode = KeyDisplayToggle
	KeySpellCheck              KeyEventCode = 0x1b0
	KeyLogOff                  KeyEventCode = 0x1b1
	KeyDollar                  KeyEventCode = 0x1b2
	KeyEuro                    KeyEventCode = 0x1b3
	KeyFrameBack               KeyEventCode = 0x1b4
	KeyFrameForward            KeyEventCode = 0x1b5
	KeyContextMenu             KeyEventCode = 0x1b6
	KeyMediaRepeat             KeyEventCode = 0x1b7
	Key10ChannelsUp            KeyEventCode = 0x1b8
	Key10ChannelsDown          KeyEventCode = 0x1b9
	KeyImages                  KeyEventCode = 0x1ba
	KeyNotificationCenter      KeyEventCode = 0x1bc
	KeyPickupPhone             KeyEventCode = 0x1bd
	KeyHangupPhone             KeyEventCode = 0x1be
	KeyDelEOL                  KeyEventCode = 0x1c0
	KeyDelEOS                  KeyEventCode = 0x1c1
	KeyInsLine                 KeyEventCode = 0x1c2
	KeyDelLine                 KeyEventCode = 0x1c3
	KeyFn                      KeyEventCode = 0x1d0
	KeyFnEsc                   KeyEventCode = 0x1d1
	KeyFnF1                    KeyEventCode = 0x1d2
	KeyFnF2                    KeyEventCode = 0x1d3
	KeyFnF3                    KeyEventCode = 0x1d4
	KeyFnF4                    KeyEventCode = 0x1d5
	KeyFnF5                    KeyEventCode = 0x1d6
	KeyFnF6                    KeyEventCode = 0x1d7
	KeyFnF7                    KeyEventCode = 0x1d8
	KeyFnF8                    KeyEventCode = 0x1d9
	KeyFnF9                    KeyEventCode = 0x1da
	KeyFnF10                   KeyEventCode = 0x1db
	KeyFnF11                   KeyEventCode = 0x1dc
	KeyFnF12                   KeyEventCode = 0x1dd
	KeyFn1                     KeyEventCode = 0x1de
	KeyFn2                     KeyEventCode = 0x1df
	KeyFnD                     KeyEventCode = 0x1e0
	KeyFnE                     KeyEventCode = 0x1e1
	KeyFnF                     KeyEventCode = 0x1e2
	KeyFnS                     KeyEventCode = 0x1e3
	KeyFnB                     KeyEventCode = 0x1e4
	KeyFnRightShift            KeyEventCode = 0x1e5
	KeyBrlDot1                 KeyEventCode = 0x1f1
	KeyBrlDot2                 KeyEventCode = 0x1f2
	KeyBrlDot3                 KeyEventCode = 0x1f3
	KeyBrlDot4                 KeyEventCode = 0x1f4
	KeyBrlDot5                 KeyEventCode = 0x1f5
	KeyBrlDot6                 KeyEventCode = 0x1f6
	KeyBrlDot7                 KeyEventCode = 0x1f7
	KeyBrlDot8                 KeyEventCode = 0x1f8
	KeyBrlDot9                 KeyEventCode = 0x1f9
	KeyBrlDot10                KeyEventCode = 0x1fa
	KeyNumeric0                KeyEventCode = 0x200
	KeyNumeric1                KeyEventCode = 0x201
	KeyNumeric2                KeyEventCode = 0x202
	KeyNumeric3                KeyEventCode = 0x203
	KeyNumeric4                KeyEventCode = 0x204
	KeyNumeric5                KeyEventCode = 0x205
	KeyNumeric6                KeyEventCode = 0x206
	KeyNumeric7                KeyEventCode = 0x207
	KeyNumeric8                KeyEventCode = 0x208
	KeyNumeric9                KeyEventCode = 0x209
	KeyNumericStar             KeyEventCode = 0x20a
	KeyNumericPound            KeyEventCode = 0x20b
	KeyNumericA                KeyEventCode = 0x20c
	KeyNumericB                KeyEventCode = 0x20d
	KeyNumericC                KeyEventCode = 0x20e
	KeyNumericD                KeyEventCode = 0x20f
	KeyCameraFocus             KeyEventCode = 0x210
	KeyWPSButton               KeyEventCode = 0x211
	KeyTouchpadToggle          KeyEventCode = 0x212
	KeyTouchpadOn              KeyEventCode = 0x213
	KeyTouchpadOff             KeyEventCode = 0x214
	KeyCameraZoomIn            KeyEventCode = 0x215
	KeyCameraZoomOut           KeyEventCode = 0x216
	KeyCameraUp                KeyEventCode = 0x217
	KeyCameraDown              KeyEventCode = 0x218
	KeyCameraLeft              KeyEventCode = 0x219
	KeyCameraRight             KeyEventCode = 0x21a
	KeyAttendantOn             KeyEventCode = 0x21b
	KeyAttendantOff            KeyEventCode = 0x21c
	KeyAttendantToggle         KeyEventCode = 0x21d
	KeyLightsToggle            KeyEventCode = 0x21e
	KeyALS_Toggle              KeyEventCode = 0x230
	KeyRotateLockToggle        KeyEventCode = 0x231
	KeyRefreshRateToggle       KeyEventCode = 0x232
	KeyButtonConfig            KeyEventCode = 0x240
	KeyTaskManager             KeyEventCode = 0x241
	KeyJournal                 KeyEventCode = 0x242
	KeyControlPanel            KeyEventCode = 0x243
	KeyAppSelect               KeyEventCode = 0x244
	KeyScreenSaver             KeyEventCode = 0x245
	KeyVoiceCommand            KeyEventCode = 0x246
	KeyAssistant               KeyEventCode = 0x247
	KeyKbdLayoutNext           KeyEventCode = 0x248
	KeyEmojiPicker             KeyEventCode = 0x249
	KeyDictate                 KeyEventCode = 0x24a
	KeyCameraAccessEnable      KeyEventCode = 0x24b
	KeyCameraAccessDisable     KeyEventCode = 0x24c
	KeyCameraAccessToggle      KeyEventCode = 0x24d
	KeyAccessibility           KeyEventCode = 0x24e
	KeyDoNotDisturb            KeyEventCode = 0x24f
	KeyBrightnessMin           KeyEventCode = 0x250
	KeyBrightnessMax           KeyEventCode = 0x251
	KeyKbdInputAssistPrev      KeyEventCode = 0x260
	KeyKbdInputAssistNext      KeyEventCode = 0x261
	KeyKbdInputAssistPrevGroup KeyEventCode = 0x262
	KeyKbdInputAssistNextGroup KeyEventCode = 0x263
	KeyKbdInputAssistAccept    KeyEventCode = 0x264
	KeyKbdInputAssistCancel    KeyEventCode = 0x265
	KeyRightUp                 KeyEventCode = 0x266
	KeyRightDown               KeyEventCode = 0x267
	KeyLeftUp                  KeyEventCode = 0x268
	KeyLeftDown                KeyEventCode = 0x269
	KeyRootMenu                KeyEventCode = 0x26a
	KeyMediaTopMenu            KeyEventCode = 0x26b
	KeyNumeric11               KeyEventCode = 0x26c
	KeyNumeric12               KeyEventCode = 0x26d
	KeyAudioDesc               KeyEventCode = 0x26e
	Key3DMode                  KeyEventCode = 0x26f
	KeyNextFavorite            KeyEventCode = 0x270
	KeyStopRecord              KeyEventCode = 0x271
	KeyPauseRecord             KeyEventCode = 0x272
	KeyVOD                     KeyEventCode = 0x273
	KeyUnmute                  KeyEventCode = 0x274
	KeyFastReverse             KeyEventCode = 0x275
	KeySlowReverse             KeyEventCode = 0x276
	KeyData                    KeyEventCode = 0x277
	KeyOnScreenKeyboard        KeyEventCode = 0x278
	KeyPrivacyScreenToggle     KeyEventCode = 0x279
	KeySelectiveScreenshot     KeyEventCode = 0x27a
	KeyNextElement             KeyEventCode = 0x27b
	KeyPreviousElement         KeyEventCode = 0x27c
	KeyAutopilotEngageToggle   KeyEventCode = 0x27d
	KeyMarkWaypoint            KeyEventCode = 0x27e
	KeySOS                     KeyEventCode = 0x27f
	KeyNavChart                KeyEventCode = 0x280
	KeyFishingChart            KeyEventCode = 0x281
	KeySingleRangeRadar        KeyEventCode = 0x282
	KeyDualRangeRadar          KeyEventCode = 0x283
	KeyRadarOverlay            KeyEventCode = 0x284
	KeyTraditionalSonar        KeyEventCode = 0x285
	KeyClearVUSonar            KeyEventCode = 0x286
	KeySideVUSonar             KeyEventCode = 0x287
	KeyNavInfo                 KeyEventCode = 0x288
	KeyBrightnessMenu          KeyEventCode = 0x289
	KeyMacro1                  KeyEventCode = 0x290
	KeyMacro2                  KeyEventCode = 0x291
	KeyMacro3                  KeyEventCode = 0x292
	KeyMacro4                  KeyEventCode = 0x293
	KeyMacro5                  KeyEventCode = 0x294
	KeyMacro6                  KeyEventCode = 0x295
	KeyMacro7                  KeyEventCode = 0x296
	KeyMacro8                  KeyEventCode = 0x297
	KeyMacro9                  KeyEventCode = 0x298
	KeyMacro10                 KeyEventCode = 0x299
	KeyMacro11                 KeyEventCode = 0x29a
	KeyMacro12                 KeyEventCode = 0x29b
	KeyMacro13                 KeyEventCode = 0x29c
	KeyMacro14                 KeyEventCode = 0x29d
	KeyMacro15                 KeyEventCode = 0x29e
	KeyMacro16                 KeyEventCode = 0x29f
	KeyMacro17                 KeyEventCode = 0x2a0
	KeyMacro18                 KeyEventCode = 0x2a1
	KeyMacro19                 KeyEventCode = 0x2a2
	KeyMacro20                 KeyEventCode = 0x2a3
	KeyMacro21                 KeyEventCode = 0x2a4
	KeyMacro22                 KeyEventCode = 0x2a5
	KeyMacro23                 KeyEventCode = 0x2a6
	KeyMacro24                 KeyEventCode = 0x2a7
	KeyMacro25                 KeyEventCode = 0x2a8
	KeyMacro26                 KeyEventCode = 0x2a9
	KeyMacro27                 KeyEventCode = 0x2aa
	KeyMacro28                 KeyEventCode = 0x2ab
	KeyMacro29                 KeyEventCode = 0x2ac
	KeyMacro30                 KeyEventCode = 0x2ad
	KeyMacroRecordStart        KeyEventCode = 0x2b0
	KeyMacroRecordStop         KeyEventCode = 0x2b1
	KeyMacroPresetCycle        KeyEventCode = 0x2b2
	KeyMacroPreset1            KeyEventCode = 0x2b3
	KeyMacroPreset2            KeyEventCode = 0x2b4
	KeyMacroPreset3            KeyEventCode = 0x2b5
	KeyKbdLcdMenu1             KeyEventCode = 0x2b8
	KeyKbdLcdMenu2             KeyEventCode = 0x2b9
	KeyKbdLcdMenu3             KeyEventCode = 0x2ba
	KeyKbdLcdMenu4             KeyEventCode = 0x2bb
	KeyKbdLcdMenu5             KeyEventCode = 0x2bc
	BtnTriggerHappy            KeyEventCode = 0x2c0
	BtnTriggerHappy1           KeyEventCode = 0x2c0
	BtnTriggerHappy2           KeyEventCode = 0x2c1
	BtnTriggerHappy3           KeyEventCode = 0x2c2
	BtnTriggerHappy4           KeyEventCode = 0x2c3
	BtnTriggerHappy5           KeyEventCode = 0x2c4
	BtnTriggerHappy6           KeyEventCode = 0x2c5
	BtnTriggerHappy7           KeyEventCode = 0x2c6
	BtnTriggerHappy8           KeyEventCode = 0x2c7
	BtnTriggerHappy9           KeyEventCode = 0x2c8
	BtnTriggerHappy10          KeyEventCode = 0x2c9
	BtnTriggerHappy11          KeyEventCode = 0x2ca
	BtnTriggerHappy12          KeyEventCode = 0x2cb
	BtnTriggerHappy13          KeyEventCode = 0x2cc
	BtnTriggerHappy14          KeyEventCode = 0x2cd
	BtnTriggerHappy15          KeyEventCode = 0x2ce
	BtnTriggerHappy16          KeyEventCode = 0x2cf
	BtnTriggerHappy17          KeyEventCode = 0x2d0
	BtnTriggerHappy18          KeyEventCode = 0x2d1
	BtnTriggerHappy19          KeyEventCode = 0x2d2
	BtnTriggerHappy20          KeyEventCode = 0x2d3
	BtnTriggerHappy21          KeyEventCode = 0x2d4
	BtnTriggerHappy22          KeyEventCode = 0x2d5
	BtnTriggerHappy23          KeyEventCode = 0x2d6
	BtnTriggerHappy24          KeyEventCode = 0x2d7
	BtnTriggerHappy25          KeyEventCode = 0x2d8
	BtnTriggerHappy26          KeyEventCode = 0x2d9
	BtnTriggerHappy27          KeyEventCode = 0x2da
	BtnTriggerHappy28          KeyEventCode = 0x2db
	BtnTriggerHappy29          KeyEventCode = 0x2dc
	BtnTriggerHappy30          KeyEventCode = 0x2dd
	BtnTriggerHappy31          KeyEventCode = 0x2de
	BtnTriggerHappy32          KeyEventCode = 0x2df
	BtnTriggerHappy33          KeyEventCode = 0x2e0
	BtnTriggerHappy34          KeyEventCode = 0x2e1
	BtnTriggerHappy35          KeyEventCode = 0x2e2
	BtnTriggerHappy36          KeyEventCode = 0x2e3
	BtnTriggerHappy37          KeyEventCode = 0x2e4
	BtnTriggerHappy38          KeyEventCode = 0x2e5
	BtnTriggerHappy39          KeyEventCode = 0x2e6
	BtnTriggerHappy40          KeyEventCode = 0x2e7

	KeyMinInteresting KeyEventCode = KeyMute
	KeyMax            KeyEventCode = 0x2ff
	KeyCnt            KeyEventCode = KeyMax + 1
)

//go:generate stringer -type=RelativeEventCode
type RelativeEventCode uint16

func (c RelativeEventCode) Type() string {
	return "Mouse"
}

func (c RelativeEventCode) Value() uint16 {
	return uint16(c)
}

const (
	RelX      RelativeEventCode = 0
	RelY      RelativeEventCode = 1
	RelZ      RelativeEventCode = 2
	RelRx     RelativeEventCode = 3
	RelRy     RelativeEventCode = 4
	RelRz     RelativeEventCode = 5
	RelHWheel RelativeEventCode = 6
	RelDial   RelativeEventCode = 7
	RelWheel  RelativeEventCode = 8
	RelMisc   RelativeEventCode = 9

	RelReserved    RelativeEventCode = 0x0a
	RelWheelHiRes  RelativeEventCode = 0x0b
	RelHWheelHiRes RelativeEventCode = 0x0c
	RelMax         RelativeEventCode = 0x0f
	RelCnt         RelativeEventCode = RelMax + 1
)

type AbsoluteAxesEventCode uint16

func (c AbsoluteAxesEventCode) Type() string {
	return "Absolute axes"
}

func (c AbsoluteAxesEventCode) Value() uint16 {
	return uint16(c)
}

const (
	AbsX         AbsoluteAxesEventCode = 0
	AbsY         AbsoluteAxesEventCode = 1
	AbsZ         AbsoluteAxesEventCode = 2
	AbsRx        AbsoluteAxesEventCode = 3
	AbsRy        AbsoluteAxesEventCode = 4
	AbsRz        AbsoluteAxesEventCode = 5
	AbsThrottle  AbsoluteAxesEventCode = 6
	AbsRudder    AbsoluteAxesEventCode = 7
	AbsWheel     AbsoluteAxesEventCode = 8
	AbsGas       AbsoluteAxesEventCode = 9
	AbsBrake     AbsoluteAxesEventCode = 10
	AbsHat0X     AbsoluteAxesEventCode = 0x10
	AbsHat0Y     AbsoluteAxesEventCode = 0x11
	AbsHat1X     AbsoluteAxesEventCode = 0x12
	AbsHat1Y     AbsoluteAxesEventCode = 0x13
	AbsHat2X     AbsoluteAxesEventCode = 0x14
	AbsHat2Y     AbsoluteAxesEventCode = 0x15
	AbsHat3X     AbsoluteAxesEventCode = 0x16
	AbsHat3Y     AbsoluteAxesEventCode = 0x17
	AbsPressure  AbsoluteAxesEventCode = 0x18
	AbsDistance  AbsoluteAxesEventCode = 0x19
	AbsTiltX     AbsoluteAxesEventCode = 0x1a
	AbsTiltY     AbsoluteAxesEventCode = 0x1b
	AbsToolWidth AbsoluteAxesEventCode = 0x1c
	AbsVolume    AbsoluteAxesEventCode = 0x20
	AbsMisc      AbsoluteAxesEventCode = 0x28
	AbsReserved  AbsoluteAxesEventCode = 0x2e

	AbsMtSlot        AbsoluteAxesEventCode = 0x2f
	AbsMtTouchMajor  AbsoluteAxesEventCode = 0x30
	AbsMtTouchMinor  AbsoluteAxesEventCode = 0x31
	AbsMtWidthMajor  AbsoluteAxesEventCode = 0x32
	AbsMtWidthMinor  AbsoluteAxesEventCode = 0x33
	AbsMtOrientation AbsoluteAxesEventCode = 0x34
	AbsMtPositionX   AbsoluteAxesEventCode = 0x35
	AbsMtPositionY   AbsoluteAxesEventCode = 0x36
	AbsMtToolType    AbsoluteAxesEventCode = 0x37
	AbsMtBlobId      AbsoluteAxesEventCode = 0x38
	AbsMtTrackingId  AbsoluteAxesEventCode = 0x39
	AbsMtPressure    AbsoluteAxesEventCode = 0x3a
	AbsMtDistance    AbsoluteAxesEventCode = 0x3b
	AbsMtToolX       AbsoluteAxesEventCode = 0x3c
	AbsMtToolY       AbsoluteAxesEventCode = 0x3d

	AbsMax AbsoluteAxesEventCode = 0x3f
	AbsCnt AbsoluteAxesEventCode = AbsMax + 1
)

//go:generate stringer -type=SyncEventCode
type SyncEventCode uint16

func (c SyncEventCode) Type() string {
	return "Sync"
}

func (c SyncEventCode) Value() uint16 {
	return uint16(c)
}

const (
	SynReport   SyncEventCode = 0
	SynConfig   SyncEventCode = 1
	SynMTReport SyncEventCode = 2
	SynDropped  SyncEventCode = 3
	SynMax      SyncEventCode = 0xf
	SynCnt      SyncEventCode = SynMax + 1
)

//go:generate stringer -type=MiscEventCode
type MiscEventCode uint16

func (c MiscEventCode) Type() string {
	return "Misc"
}

func (c MiscEventCode) Value() uint16 {
	return uint16(c)
}

const (
	MscSerial    MiscEventCode = 0
	MscPulseLED  MiscEventCode = 1
	MscGesture   MiscEventCode = 2
	MscRaw       MiscEventCode = 3
	MscScan      MiscEventCode = 4
	MscTimestamp MiscEventCode = 5
	MscMax       MiscEventCode = 0x7
	MscCnt       MiscEventCode = MscMax + 1
)

//go:generate stringer -type=SwitchEventCode
type SwitchEventCode uint16

func (c SwitchEventCode) Type() string {
	return "Switch"
}

func (c SwitchEventCode) Value() uint16 {
	return uint16(c)
}

const (
	SwLid                SwitchEventCode = 0x00
	SwTabletMode         SwitchEventCode = 0x01
	SwHeadphoneInsert    SwitchEventCode = 0x02
	SwRfkillAll          SwitchEventCode = 0x03
	SwMicroPhoneInsert   SwitchEventCode = 0x04
	SwDock               SwitchEventCode = 0x05
	SwLineoutInsert      SwitchEventCode = 0x06
	SwJackPhysicalInsert SwitchEventCode = 0x07
	SwVideoOutInsert     SwitchEventCode = 0x08
	SwCameraLensCover    SwitchEventCode = 0x09
	SwKeypadSlide        SwitchEventCode = 0x0a
	SwFrontProximity     SwitchEventCode = 0x0b
	SwRotateLock         SwitchEventCode = 0x0c
	SwLineinInsert       SwitchEventCode = 0x0d
	SwMuteDevice         SwitchEventCode = 0x0e
	SwPenInsert          SwitchEventCode = 0x0f
	SwMax                SwitchEventCode = 0x0f
	SwCnt                SwitchEventCode = SwMax + 1
)

//go:generate stringer -type=LEDEventCode
type LEDEventCode uint16

func (c LEDEventCode) Type() string {
	return "LED"
}

func (c LEDEventCode) Value() uint16 {
	return uint16(c)
}

const (
	LedNumLock    LEDEventCode = 0x00
	LedCapsLock   LEDEventCode = 0x01
	LedScrollLock LEDEventCode = 0x02
	LedCompose    LEDEventCode = 0x03
	LedKana       LEDEventCode = 0x04
	LedSleep      LEDEventCode = 0x05
	LedSuspend    LEDEventCode = 0x06
	LedMute       LEDEventCode = 0x07
	LedMisc       LEDEventCode = 0x08
	LedMail       LEDEventCode = 0x09
	LedCharging   LEDEventCode = 0x0a
	LedMax        LEDEventCode = 0x0f
	LedCnt        LEDEventCode = LedMax + 1
)

//go:generate stringer -type=SoundEventCode
type SoundEventCode uint16

func (c SoundEventCode) Type() string {
	return "Sound"
}

func (c SoundEventCode) Value() uint16 {
	return uint16(c)
}

const (
	SndClick SoundEventCode = 0x00
	SndBell  SoundEventCode = 0x01
	SndTone  SoundEventCode = 0x02
	SndMax   SoundEventCode = 0x07
	SndCnt   SoundEventCode = SndMax + 1
)

//go:generate stringer -type=RepeatEventCode
type RepeatEventCode uint16

func (c RepeatEventCode) Type() string {
	return "Repeat"
}

func (c RepeatEventCode) Value() uint16 {
	return uint16(c)
}

const (
	RepDelay  RepeatEventCode = 0x00
	RepPeriod RepeatEventCode = 0x01
	RepMax    RepeatEventCode = 0x01
	RepCnt    RepeatEventCode = RepMax + 1
)

//go:generate stringer -type=ForceEventCode
type ForceEventCode uint16

func (c ForceEventCode) Type() string {
	return "Force"
}

func (c ForceEventCode) Value() uint16 {
	return uint16(c)
}

const (
	FfRumble       ForceEventCode = 0x50
	FfPeriodic     ForceEventCode = 0x51
	FfConstant     ForceEventCode = 0x52
	FfSpring       ForceEventCode = 0x53
	FfFriction     ForceEventCode = 0x54
	FfDamper       ForceEventCode = 0x55
	FfInertia      ForceEventCode = 0x56
	FfRamp         ForceEventCode = 0x57
	FfSawtoothUp   ForceEventCode = 0x58
	FfSawtoothDown ForceEventCode = 0x59
	FfCustom       ForceEventCode = 0x5a
	FfGain         ForceEventCode = 0x60
	FfAutoCenter   ForceEventCode = 0x61
	FfMax          ForceEventCode = 0x7f
	FfCnt          ForceEventCode = FfMax + 1
)

//go:generate stringer -type=FFStatusEventCode
type FFStatusEventCode uint16

func (c FFStatusEventCode) Type() string {
	return "Force Feedback Status"
}

func (c FFStatusEventCode) Value() uint16 {
	return uint16(c)
}

const (
	FfStatusStopped FFStatusEventCode = 0x00
	FfStatusPlaying FFStatusEventCode = 0x01
	FfStatusMax     FFStatusEventCode = 0x01
	FfStatusCnt     FFStatusEventCode = FfStatusMax + 1
)

//go:generate stringer -type=PowerEventCode
type PowerEventCode uint16

func (c PowerEventCode) Type() string {
	return "Power"
}

func (c PowerEventCode) Value() uint16 {
	return uint16(c)
}

const (
	PwrName     PowerEventCode = 0x00
	PwrSerial   PowerEventCode = 0x01
	PwrCapacity PowerEventCode = 0x02
	PwrStatus   PowerEventCode = 0x03
	PwrPresent  PowerEventCode = 0x04
	PwrReserved PowerEventCode = 0x05
	PwrCycle    PowerEventCode = 0x06
	PwrMax      PowerEventCode = 0x0f
	PwrCnt      PowerEventCode = PwrMax + 1
)

type UnknownEventCode uint16

func (c UnknownEventCode) Type() string {
	return "Unknown"
}

func (c UnknownEventCode) Value() uint16 {
	return uint16(c)
}
