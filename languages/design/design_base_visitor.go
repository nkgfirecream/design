// Code generated from Design.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Design

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDesignVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDesignVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigDeclaration(ctx *ConfigDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigKey(ctx *ConfigKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitConfigValue(ctx *ConfigValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDecalartions(ctx *DecalartionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitFlowDeclaration(ctx *FlowDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitInteractionDeclaration(ctx *InteractionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitSeeDeclaration(ctx *SeeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitDoDeclaration(ctx *DoDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitReactDeclaration(ctx *ReactDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitAnimateDeclaration(ctx *AnimateDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitReactAction(ctx *ReactActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitGotoAction(ctx *GotoActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitShowAction(ctx *ShowActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitActionName(ctx *ActionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentValue(ctx *ComponentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentName(ctx *ComponentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitSceneName(ctx *SceneNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitAnimateName(ctx *AnimateNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPageDeclaration(ctx *PageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentDeclaration(ctx *ComponentDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentBodyDeclaration(ctx *ComponentBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutDeclaration(ctx *LayoutDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutRow(ctx *LayoutRowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutLines(ctx *LayoutLinesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLayoutLine(ctx *LayoutLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentUseDeclaration(ctx *ComponentUseDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitComponentLayoutValue(ctx *ComponentLayoutValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleDeclaration(ctx *StyleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleName(ctx *StyleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitStyleBody(ctx *StyleBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryDeclaration(ctx *LibraryDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryExpress(ctx *LibraryExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPresetKey(ctx *PresetKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPresetValue(ctx *PresetValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPresetArray(ctx *PresetArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitPresetCall(ctx *PresetCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDesignVisitor) VisitLibraryName(ctx *LibraryNameContext) interface{} {
	return v.VisitChildren(ctx)
}
