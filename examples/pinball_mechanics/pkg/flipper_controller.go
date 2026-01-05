package pkg

import (
	"strings"

	. "github.com/godot-go/godot-go/pkg/builtin"
	. "github.com/godot-go/godot-go/pkg/constant"
	. "github.com/godot-go/godot-go/pkg/core"
	. "github.com/godot-go/godot-go/pkg/gdclassimpl"
	"github.com/godot-go/godot-go/pkg/log"
)

// FlipperController implements GDClass evidence.
var _ GDClass = (*FlipperController)(nil)

type FlipperController struct {
	Node2DImpl
	keycode     Key
	initialized bool
	motorSpeed  float32
	joint       PinJoint2D
	flipperBody RigidBody2D
}

func (f *FlipperController) GetClassName() string {
	return "FlipperController"
}

func (f *FlipperController) GetParentClassName() string {
	return "Node2D"
}

func (f *FlipperController) V_Ready() {
	log.Info("FlipperController: _ready")
	f.SetPhysicsProcess(true)
	f.configureSide()
	f.cacheNodes()
}

func (f *FlipperController) V_PhysicsProcess(_delta float64) {
	if !f.initialized || f.joint == nil {
		return
	}
	input := inputSingleton()
	if input == nil {
		return
	}
	if input.IsKeyPressed(f.keycode) {
		f.joint.SetMotorTargetVelocity(f.motorSpeed)
	} else {
		f.joint.SetMotorTargetVelocity(0)
	}
}

func (f *FlipperController) V_ExitTree() {
	f.initialized = false
}

func (f *FlipperController) configureSide() {
	name := f.GetName()
	isLeft := strings.Contains(strings.ToLower(name.ToUtf8()), "left")
	name.Destroy()
	if isLeft {
		f.keycode = KEY_LEFT
		f.motorSpeed = 12.0
	} else {
		f.keycode = KEY_RIGHT
		f.motorSpeed = -12.0
	}
	f.initialized = true
}

func (f *FlipperController) cacheNodes() {
	bodyPath := nodePath("FlipperBody")
	defer bodyPath.Destroy()
	jointPath := nodePath("FlipperJoint")
	defer jointPath.Destroy()
	anchorPath := nodePath("FlipperAnchor")
	defer anchorPath.Destroy()

	bodyNode := f.GetNodeOrNull(bodyPath)
	jointNode := f.GetNodeOrNull(jointPath)
	anchorNode := f.GetNodeOrNull(anchorPath)
	if bodyNode == nil || jointNode == nil || anchorNode == nil {
		printLine("FlipperController: missing child nodes")
		return
	}

	f.flipperBody, _ = ObjectCastTo(bodyNode, "RigidBody2D").(RigidBody2D)
	f.joint, _ = ObjectCastTo(jointNode, "PinJoint2D").(PinJoint2D)
	anchor, _ := ObjectCastTo(anchorNode, "StaticBody2D").(StaticBody2D)
	if f.flipperBody == nil || f.joint == nil || anchor == nil {
		printLine("FlipperController: child cast failed")
		return
	}

}

func (f *FlipperController) configureJoint() {
	if f.joint == nil {
		return
	}
	f.joint.SetMotorEnabled(true)
	f.joint.SetAngularLimitEnabled(true)
	f.joint.SetSoftness(0.2)
	if f.motorSpeed > 0 {
		f.joint.SetAngularLimitLower(-1.0)
		f.joint.SetAngularLimitUpper(0.2)
	} else {
		f.joint.SetAngularLimitLower(-0.2)
		f.joint.SetAngularLimitUpper(1.0)
	}
}

func NewFlipperControllerFromOwnerObject(owner *GodotObject) GDClass {
	obj := &FlipperController{}
	obj.SetGodotObjectOwner(owner)
	return obj
}

func RegisterClassFlipperController() {
	ClassDBRegisterClass(NewFlipperControllerFromOwnerObject, nil, nil, func(t *FlipperController) {
		ClassDBBindMethodVirtual(t, "V_Ready", "_ready", nil, nil)
		ClassDBBindMethodVirtual(t, "V_PhysicsProcess", "_physics_process", []string{"delta"}, nil)
		ClassDBBindMethodVirtual(t, "V_ExitTree", "_exit_tree", nil, nil)
	})
}

func UnregisterClassFlipperController() {
	ClassDBUnregisterClass[*FlipperController]()
}
