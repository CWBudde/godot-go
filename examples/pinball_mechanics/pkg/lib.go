package pkg

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../../../godot_headers -I${SRCDIR}/../../../pkg/log -I${SRCDIR}/../../../pkg/gdextension
*/
import "C"

import (
	"unsafe"

	. "github.com/godot-go/godot-go/pkg/core"
	"github.com/godot-go/godot-go/pkg/ffi"
	"github.com/godot-go/godot-go/pkg/log"
	"github.com/godot-go/godot-go/pkg/util"
)

func registerPinballTypes() {
	log.Info("pinball: registering PinballGame")
	RegisterClassPinballGame()
	log.Info("pinball: registering FlipperController")
	RegisterClassFlipperController()
	log.Info("pinball: registering BallLauncher")
	RegisterClassBallLauncher()
	log.Info("pinball: registering Bumper")
	RegisterClassBumper()
	log.Info("pinball: registering DrainDetector")
	RegisterClassDrainDetector()
	log.Info("pinball: registering ScoreSystem")
	RegisterClassScoreSystem()
}

func unregisterPinballTypes() {
	UnregisterClassPinballGame()
	UnregisterClassFlipperController()
	UnregisterClassBallLauncher()
	UnregisterClassBumper()
	UnregisterClassDrainDetector()
	UnregisterClassScoreSystem()
}

//export PinballInit
func PinballInit(pGetProcAddress unsafe.Pointer, pLibrary unsafe.Pointer, rInitialization unsafe.Pointer) bool {
	util.SetThreadName("pinball-mechanics")
	log.Debug("PinballInit called")
	initObj := NewInitObject(
		(ffi.GDExtensionInterfaceGetProcAddress)(pGetProcAddress),
		(ffi.GDExtensionClassLibraryPtr)(pLibrary),
		(*ffi.GDExtensionInitialization)(rInitialization),
	)
	initObj.RegisterSceneInitializer(registerPinballTypes)
	initObj.RegisterSceneTerminator(unregisterPinballTypes)
	return initObj.Init()
}
