// SPDX-License-Identifier: Unlicense OR MIT

//go:build windows
// +build windows

package windows

import (
	"fmt"
	"runtime"
	"structs"
	"time"
	"unicode/utf16"
	"unsafe"

	syscall "golang.org/x/sys/windows"
)

type CompositionForm struct {
	dwStyle      uint32
	ptCurrentPos Point
	rcArea       Rect
}

type CandidateForm struct {
	dwIndex      uint32
	dwStyle      uint32
	ptCurrentPos Point
	rcArea       Rect
}

type Rect struct {
	Left, Top, Right, Bottom int32
}

type WndClassEx struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CnClsExtra    int32
	CbWndExtra    int32
	HInstance     syscall.Handle
	HIcon         syscall.Handle
	HCursor       syscall.Handle
	HbrBackground syscall.Handle
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       syscall.Handle
}

type Margins struct {
	CxLeftWidth    int32
	CxRightWidth   int32
	CyTopHeight    int32
	CyBottomHeight int32
}

type Msg struct {
	Hwnd     syscall.Handle
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       Point
	LPrivate uint32
}

type Point struct {
	X, Y int32
}

type MinMaxInfo struct {
	PtReserved     Point
	PtMaxSize      Point
	PtMaxPosition  Point
	PtMinTrackSize Point
	PtMaxTrackSize Point
}

type NCCalcSizeParams struct {
	Rgrc  [3]Rect
	LpPos *WindowPos
}

type WindowPos struct {
	HWND            syscall.Handle
	HWNDInsertAfter syscall.Handle
	x               int32
	y               int32
	cx              int32
	cy              int32
	flags           uint32
}

type WindowPlacement struct {
	length           uint32
	flags            uint32
	showCmd          uint32
	ptMinPosition    Point
	ptMaxPosition    Point
	rcNormalPosition Rect
	rcDevice         Rect
}

type MonitorInfo struct {
	cbSize   uint32
	Monitor  Rect
	WorkArea Rect
	Flags    uint32
}

type CopyDataStruct struct {
	DwData uintptr
	CbData uint32
	LpData uintptr
}

type POINTER_INPUT_TYPE int32

const (
	PT_POINTER  POINTER_INPUT_TYPE = 1
	PT_TOUCH    POINTER_INPUT_TYPE = 2
	PT_PEN      POINTER_INPUT_TYPE = 3
	PT_MOUSE    POINTER_INPUT_TYPE = 4
	PT_TOUCHPAD POINTER_INPUT_TYPE = 5
)

type POINTER_INFO_POINTER_FLAGS int32

const (
	POINTER_FLAG_NEW            POINTER_INFO_POINTER_FLAGS = 0x00000001
	POINTER_FLAG_INRANGE        POINTER_INFO_POINTER_FLAGS = 0x00000002
	POINTER_FLAG_INCONTACT      POINTER_INFO_POINTER_FLAGS = 0x00000004
	POINTER_FLAG_FIRSTBUTTON    POINTER_INFO_POINTER_FLAGS = 0x00000010
	POINTER_FLAG_SECONDBUTTON   POINTER_INFO_POINTER_FLAGS = 0x00000020
	POINTER_FLAG_THIRDBUTTON    POINTER_INFO_POINTER_FLAGS = 0x00000040
	POINTER_FLAG_FOURTHBUTTON   POINTER_INFO_POINTER_FLAGS = 0x00000080
	POINTER_FLAG_FIFTHBUTTON    POINTER_INFO_POINTER_FLAGS = 0x00000100
	POINTER_FLAG_PRIMARY        POINTER_INFO_POINTER_FLAGS = 0x00002000
	POINTER_FLAG_CONFIDENCE     POINTER_INFO_POINTER_FLAGS = 0x00004000
	POINTER_FLAG_CANCELED       POINTER_INFO_POINTER_FLAGS = 0x00008000
	POINTER_FLAG_DOWN           POINTER_INFO_POINTER_FLAGS = 0x00010000
	POINTER_FLAG_UPDATE         POINTER_INFO_POINTER_FLAGS = 0x00020000
	POINTER_FLAG_UP             POINTER_INFO_POINTER_FLAGS = 0x00040000
	POINTER_FLAG_WHEEL          POINTER_INFO_POINTER_FLAGS = 0x00080000
	POINTER_FLAG_HWHEEL         POINTER_INFO_POINTER_FLAGS = 0x00100000
	POINTER_FLAG_CAPTURECHANGED POINTER_INFO_POINTER_FLAGS = 0x00200000
	POINTER_FLAG_HASTRANSFORM   POINTER_INFO_POINTER_FLAGS = 0x00400000
)

type POINTER_BUTTON_CHANGE_TYPE int32

const (
	POINTER_CHANGE_NONE              POINTER_BUTTON_CHANGE_TYPE = 0
	POINTER_CHANGE_FIRSTBUTTON_DOWN  POINTER_BUTTON_CHANGE_TYPE = 1
	POINTER_CHANGE_FIRSTBUTTON_UP    POINTER_BUTTON_CHANGE_TYPE = 2
	POINTER_CHANGE_SECONDBUTTON_DOWN POINTER_BUTTON_CHANGE_TYPE = 3
	POINTER_CHANGE_SECONDBUTTON_UP   POINTER_BUTTON_CHANGE_TYPE = 4
	POINTER_CHANGE_THIRDBUTTON_DOWN  POINTER_BUTTON_CHANGE_TYPE = 5
	POINTER_CHANGE_THIRDBUTTON_UP    POINTER_BUTTON_CHANGE_TYPE = 6
	POINTER_CHANGE_FOURTHBUTTON_DOWN POINTER_BUTTON_CHANGE_TYPE = 7
	POINTER_CHANGE_FOURTHBUTTON_UP   POINTER_BUTTON_CHANGE_TYPE = 8
	POINTER_CHANGE_FIFTHBUTTON_DOWN  POINTER_BUTTON_CHANGE_TYPE = 9
	POINTER_CHANGE_FIFTHBUTTON_UP    POINTER_BUTTON_CHANGE_TYPE = 10
)

type PointerInfo struct {
	PointerType           POINTER_INPUT_TYPE
	PointerId             uint32
	FrameId               uint32
	PointerFlags          POINTER_INFO_POINTER_FLAGS
	SourceDevice          syscall.Handle
	HwndTarget            syscall.Handle
	PtPixelLocation       Point
	PtHimetricLocation    Point
	PtPixelLocationRaw    Point
	PtHimetricLocationRaw Point
	DwTime                uint32
	HistoryCount          uint32
	InputData             int32
	DwKeyStates           uint32
	PerformanceCount      uint64
	ButtonChangeType      POINTER_BUTTON_CHANGE_TYPE
}

type VARIANT_TYPE uint16

const (
	VT_EMPTY            VARIANT_TYPE = 0
	VT_NULL             VARIANT_TYPE = 1
	VT_I2               VARIANT_TYPE = 2
	VT_I4               VARIANT_TYPE = 3
	VT_R4               VARIANT_TYPE = 4
	VT_R8               VARIANT_TYPE = 5
	VT_CY               VARIANT_TYPE = 6
	VT_DATE             VARIANT_TYPE = 7
	VT_BSTR             VARIANT_TYPE = 8
	VT_DISPATCH         VARIANT_TYPE = 9
	VT_ERROR            VARIANT_TYPE = 10
	VT_BOOL             VARIANT_TYPE = 11
	VT_VARIANT          VARIANT_TYPE = 12
	VT_UNKNOWN          VARIANT_TYPE = 13
	VT_DECIMAL          VARIANT_TYPE = 14
	VT_I1               VARIANT_TYPE = 16
	VT_UI1              VARIANT_TYPE = 17
	VT_UI2              VARIANT_TYPE = 18
	VT_UI4              VARIANT_TYPE = 19
	VT_I8               VARIANT_TYPE = 20
	VT_UI8              VARIANT_TYPE = 21
	VT_INT              VARIANT_TYPE = 22
	VT_UINT             VARIANT_TYPE = 23
	VT_VOID             VARIANT_TYPE = 24
	VT_HRESULT          VARIANT_TYPE = 25
	VT_PTR              VARIANT_TYPE = 26
	VT_SAFEARRAY        VARIANT_TYPE = 27
	VT_CARRAY           VARIANT_TYPE = 28
	VT_USERDEFINED      VARIANT_TYPE = 29
	VT_LPSTR            VARIANT_TYPE = 30
	VT_LPWSTR           VARIANT_TYPE = 31
	VT_RECORD           VARIANT_TYPE = 36
	VT_INT_PTR          VARIANT_TYPE = 37
	VT_UINT_PTR         VARIANT_TYPE = 38
	VT_FILETIME         VARIANT_TYPE = 64
	VT_BLOB             VARIANT_TYPE = 65
	VT_STREAM           VARIANT_TYPE = 66
	VT_STORAGE          VARIANT_TYPE = 67
	VT_STREAMED_OBJECT  VARIANT_TYPE = 68
	VT_STORED_OBJECT    VARIANT_TYPE = 69
	VT_BLOB_OBJECT      VARIANT_TYPE = 70
	VT_CF               VARIANT_TYPE = 71
	VT_CLSID            VARIANT_TYPE = 72
	VT_VERSIONED_STREAM VARIANT_TYPE = 73
	VT_BSTR_BLOB        VARIANT_TYPE = 0xfff
	VT_VECTOR           VARIANT_TYPE = 0x1000
	VT_ARRAY            VARIANT_TYPE = 0x2000
	VT_BYREF            VARIANT_TYPE = 0x4000
	VT_RESERVED         VARIANT_TYPE = 0x8000
	VT_ILLEGAL          VARIANT_TYPE = 0xffff
	VT_ILLEGALMASKED    VARIANT_TYPE = 0xfff
	VT_TYPEMASK         VARIANT_TYPE = 0xfff
)

type Variant struct {
	_ structs.HostLayout

	VT        VARIANT_TYPE
	Reserved1 uint16
	Reserved2 uint16
	Reserved3 uint16
	Val       uintptr
	_         uintptr // HACK: guarantee proper size on x64
}

type IUnknownVTable struct {
	_              structs.HostLayout
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type GUID struct {
	_       structs.HostLayout
	Data1   uint32
	Data2   uint16
	Data3   uint16
	Data4_0 uint8
	Data4_1 uint8
	Data4_2 uint8
	Data4_3 uint8
	Data4_4 uint8
	Data4_5 uint8
	Data4_6 uint8
	Data4_7 uint8
}

var (
	IID_IUnknown                        = GUID{structs.HostLayout{}, 0x00000000, 0x0000, 0x0000, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}
	IID_IUIAutomationElement            = GUID{structs.HostLayout{}, 0xd22108aa, 0x8ac5, 0x49a5, 0x83, 0x7b, 0x37, 0xbb, 0xb3, 0xd7, 0x59, 0x1e}
	IID_IUIAutomationTogglePattern      = GUID{structs.HostLayout{}, 0x94cf8058, 0x9b8d, 0x4ab9, 0x8b, 0xfd, 0x4c, 0xd0, 0xa3, 0x3c, 0x8c, 0x70}
	IID_IUIAutomationInvokePattern      = GUID{structs.HostLayout{}, 0xfb377fbe, 0x8ea6, 0x46d5, 0x9c, 0x73, 0x64, 0x99, 0x64, 0x2d, 0x30, 0x59}
	IID_IRawElementProviderSimple       = GUID{structs.HostLayout{}, 0xd6dd68d1, 0x86fd, 0x4332, 0x86, 0x66, 0x9a, 0xbe, 0xde, 0xa2, 0xd2, 0x4c}
	IID_IRawElementProviderFragment     = GUID{structs.HostLayout{}, 0xf7063da8, 0x8359, 0x439c, 0x92, 0x97, 0xbb, 0xc5, 0x29, 0x9a, 0x7d, 0x87}
	IID_IRawElementProviderFragmentRoot = GUID{structs.HostLayout{}, 0x620ce2a5, 0xab8f, 0x40a9, 0x86, 0xcb, 0xde, 0x3c, 0x75, 0x59, 0x9b, 0x58}
	IID_IInvokeProvider                 = GUID{structs.HostLayout{}, 0x54fcb24b, 0xe18e, 0x47a2, 0xb4, 0xd3, 0xec, 0xcb, 0xe7, 0x75, 0x99, 0xa2}
	IID_IValueProvider                  = GUID{structs.HostLayout{}, 0xc7935180, 0x6fb3, 0x4201, 0xb1, 0x74, 0x7d, 0xf7, 0x3a, 0xdb, 0xf6, 0x4a}
	IID_IScrollProvider                 = GUID{structs.HostLayout{}, 0xb38b8077, 0x1fc3, 0x42a5, 0x8c, 0xae, 0xd4, 0x0c, 0x22, 0x15, 0x05, 0x5a}
)

type UiaPatternId uint

const (
	UIA_AnnotationPatternId        UiaPatternId = 10023
	UIA_CustomNavigationPatternId  UiaPatternId = 10033
	UIA_DockPatternId              UiaPatternId = 10011
	UIA_DragPatternId              UiaPatternId = 10030
	UIA_DropTargetPatternId        UiaPatternId = 10031
	UIA_ExpandCollapsePatternId    UiaPatternId = 10005
	UIA_GridItemPatternId          UiaPatternId = 10007
	UIA_GridPatternId              UiaPatternId = 10006
	UIA_InvokePatternId            UiaPatternId = 10000
	UIA_ItemContainerPatternId     UiaPatternId = 10019
	UIA_LegacyIAccessiblePatternId UiaPatternId = 10018
	UIA_MultipleViewPatternId      UiaPatternId = 10008
	UIA_ObjectModelPatternId       UiaPatternId = 10022
	UIA_RangeValuePatternId        UiaPatternId = 10003
	UIA_ScrollItemPatternId        UiaPatternId = 10017
	UIA_ScrollPatternId            UiaPatternId = 10004
	UIA_SelectionItemPatternId     UiaPatternId = 10010
	UIA_SelectionPatternId         UiaPatternId = 10001
	UIA_SpreadsheetPatternId       UiaPatternId = 10026
	UIA_SpreadsheetItemPatternId   UiaPatternId = 10027
	UIA_StylesPatternId            UiaPatternId = 10025
	UIA_SynchronizedInputPatternId UiaPatternId = 10021
	UIA_TableItemPatternId         UiaPatternId = 10013
	UIA_TablePatternId             UiaPatternId = 10012
	UIA_TextChildPatternId         UiaPatternId = 10029
	UIA_TextEditPatternId          UiaPatternId = 10032
	UIA_TextPatternId              UiaPatternId = 10014
	UIA_TextPattern2Id             UiaPatternId = 10024
	UIA_TogglePatternId            UiaPatternId = 10015
	UIA_TransformPatternId         UiaPatternId = 10016
	UIA_TransformPattern2Id        UiaPatternId = 10028
	UIA_ValuePatternId             UiaPatternId = 10002
	UIA_VirtualizedItemPatternId   UiaPatternId = 10020
	UIA_WindowPatternId            UiaPatternId = 10009
)

type UiaPropertyId uint32

const (
	UIA_RuntimeIdPropertyId                           UiaPropertyId = 30000
	UIA_BoundingRectanglePropertyId                   UiaPropertyId = 30001
	UIA_ProcessIdPropertyId                           UiaPropertyId = 30002
	UIA_ControlTypePropertyId                         UiaPropertyId = 30003
	UIA_LocalizedControlTypePropertyId                UiaPropertyId = 30004
	UIA_NamePropertyId                                UiaPropertyId = 30005
	UIA_AcceleratorKeyPropertyId                      UiaPropertyId = 30006
	UIA_AccessKeyPropertyId                           UiaPropertyId = 30007
	UIA_HasKeyboardFocusPropertyId                    UiaPropertyId = 30008
	UIA_IsKeyboardFocusablePropertyId                 UiaPropertyId = 30009
	UIA_IsEnabledPropertyId                           UiaPropertyId = 30010
	UIA_AutomationIdPropertyId                        UiaPropertyId = 30011
	UIA_ClassNamePropertyId                           UiaPropertyId = 30012
	UIA_HelpTextPropertyId                            UiaPropertyId = 30013
	UIA_ClickablePointPropertyId                      UiaPropertyId = 30014
	UIA_CulturePropertyId                             UiaPropertyId = 30015
	UIA_IsControlElementPropertyId                    UiaPropertyId = 30016
	UIA_IsContentElementPropertyId                    UiaPropertyId = 30017
	UIA_LabeledByPropertyId                           UiaPropertyId = 30018
	UIA_IsPasswordPropertyId                          UiaPropertyId = 30019
	UIA_NativeWindowHandlePropertyId                  UiaPropertyId = 30020
	UIA_ItemTypePropertyId                            UiaPropertyId = 30021
	UIA_IsOffscreenPropertyId                         UiaPropertyId = 30022
	UIA_OrientationPropertyId                         UiaPropertyId = 30023
	UIA_FrameworkIdPropertyId                         UiaPropertyId = 30024
	UIA_IsRequiredForFormPropertyId                   UiaPropertyId = 30025
	UIA_ItemStatusPropertyId                          UiaPropertyId = 30026
	UIA_IsDockPatternAvailablePropertyId              UiaPropertyId = 30027
	UIA_IsExpandCollapsePatternAvailablePropertyId    UiaPropertyId = 30028
	UIA_IsGridItemPatternAvailablePropertyId          UiaPropertyId = 30029
	UIA_IsGridPatternAvailablePropertyId              UiaPropertyId = 30030
	UIA_IsInvokePatternAvailablePropertyId            UiaPropertyId = 30031
	UIA_IsMultipleViewPatternAvailablePropertyId      UiaPropertyId = 30032
	UIA_IsRangeValuePatternAvailablePropertyId        UiaPropertyId = 30033
	UIA_IsScrollPatternAvailablePropertyId            UiaPropertyId = 30034
	UIA_IsScrollItemPatternAvailablePropertyId        UiaPropertyId = 30035
	UIA_IsSelectionItemPatternAvailablePropertyId     UiaPropertyId = 30036
	UIA_IsSelectionPatternAvailablePropertyId         UiaPropertyId = 30037
	UIA_IsTablePatternAvailablePropertyId             UiaPropertyId = 30038
	UIA_IsTableItemPatternAvailablePropertyId         UiaPropertyId = 30039
	UIA_IsTextPatternAvailablePropertyId              UiaPropertyId = 30040
	UIA_IsTogglePatternAvailablePropertyId            UiaPropertyId = 30041
	UIA_IsTransformPatternAvailablePropertyId         UiaPropertyId = 30042
	UIA_IsValuePatternAvailablePropertyId             UiaPropertyId = 30043
	UIA_IsWindowPatternAvailablePropertyId            UiaPropertyId = 30044
	UIA_ValueValuePropertyId                          UiaPropertyId = 30045
	UIA_ValueIsReadOnlyPropertyId                     UiaPropertyId = 30046
	UIA_RangeValueValuePropertyId                     UiaPropertyId = 30047
	UIA_RangeValueIsReadOnlyPropertyId                UiaPropertyId = 30048
	UIA_RangeValueMinimumPropertyId                   UiaPropertyId = 30049
	UIA_RangeValueMaximumPropertyId                   UiaPropertyId = 30050
	UIA_RangeValueLargeChangePropertyId               UiaPropertyId = 30051
	UIA_RangeValueSmallChangePropertyId               UiaPropertyId = 30052
	UIA_ScrollHorizontalScrollPercentPropertyId       UiaPropertyId = 30053
	UIA_ScrollHorizontalViewSizePropertyId            UiaPropertyId = 30054
	UIA_ScrollVerticalScrollPercentPropertyId         UiaPropertyId = 30055
	UIA_ScrollVerticalViewSizePropertyId              UiaPropertyId = 30056
	UIA_ScrollHorizontallyScrollablePropertyId        UiaPropertyId = 30057
	UIA_ScrollVerticallyScrollablePropertyId          UiaPropertyId = 30058
	UIA_SelectionSelectionPropertyId                  UiaPropertyId = 30059
	UIA_SelectionCanSelectMultiplePropertyId          UiaPropertyId = 30060
	UIA_SelectionIsSelectionRequiredPropertyId        UiaPropertyId = 30061
	UIA_GridRowCountPropertyId                        UiaPropertyId = 30062
	UIA_GridColumnCountPropertyId                     UiaPropertyId = 30063
	UIA_GridItemRowPropertyId                         UiaPropertyId = 30064
	UIA_GridItemColumnPropertyId                      UiaPropertyId = 30065
	UIA_GridItemRowSpanPropertyId                     UiaPropertyId = 30066
	UIA_GridItemColumnSpanPropertyId                  UiaPropertyId = 30067
	UIA_GridItemContainingGridPropertyId              UiaPropertyId = 30068
	UIA_DockDockPositionPropertyId                    UiaPropertyId = 30069
	UIA_ExpandCollapseExpandCollapseStatePropertyId   UiaPropertyId = 30070
	UIA_MultipleViewCurrentViewPropertyId             UiaPropertyId = 30071
	UIA_MultipleViewSupportedViewsPropertyId          UiaPropertyId = 30072
	UIA_WindowCanMaximizePropertyId                   UiaPropertyId = 30073
	UIA_WindowCanMinimizePropertyId                   UiaPropertyId = 30074
	UIA_WindowWindowVisualStatePropertyId             UiaPropertyId = 30075
	UIA_WindowWindowInteractionStatePropertyId        UiaPropertyId = 30076
	UIA_WindowIsModalPropertyId                       UiaPropertyId = 30077
	UIA_WindowIsTopmostPropertyId                     UiaPropertyId = 30078
	UIA_SelectionItemIsSelectedPropertyId             UiaPropertyId = 30079
	UIA_SelectionItemSelectionContainerPropertyId     UiaPropertyId = 30080
	UIA_TableRowHeadersPropertyId                     UiaPropertyId = 30081
	UIA_TableColumnHeadersPropertyId                  UiaPropertyId = 30082
	UIA_TableRowOrColumnMajorPropertyId               UiaPropertyId = 30083
	UIA_TableItemRowHeaderItemsPropertyId             UiaPropertyId = 30084
	UIA_TableItemColumnHeaderItemsPropertyId          UiaPropertyId = 30085
	UIA_ToggleToggleStatePropertyId                   UiaPropertyId = 30086
	UIA_TransformCanMovePropertyId                    UiaPropertyId = 30087
	UIA_TransformCanResizePropertyId                  UiaPropertyId = 30088
	UIA_TransformCanRotatePropertyId                  UiaPropertyId = 30089
	UIA_IsLegacyIAccessiblePatternAvailablePropertyId UiaPropertyId = 30090
	UIA_LegacyIAccessibleChildIdPropertyId            UiaPropertyId = 30091
	UIA_LegacyIAccessibleNamePropertyId               UiaPropertyId = 30092
	UIA_LegacyIAccessibleValuePropertyId              UiaPropertyId = 30093
	UIA_LegacyIAccessibleDescriptionPropertyId        UiaPropertyId = 30094
	UIA_LegacyIAccessibleRolePropertyId               UiaPropertyId = 30095
	UIA_LegacyIAccessibleStatePropertyId              UiaPropertyId = 30096
	UIA_LegacyIAccessibleHelpPropertyId               UiaPropertyId = 30097
	UIA_LegacyIAccessibleKeyboardShortcutPropertyId   UiaPropertyId = 30098
	UIA_LegacyIAccessibleSelectionPropertyId          UiaPropertyId = 30099
	UIA_LegacyIAccessibleDefaultActionPropertyId      UiaPropertyId = 30100
	UIA_AriaRolePropertyId                            UiaPropertyId = 30101
	UIA_AriaPropertiesPropertyId                      UiaPropertyId = 30102
	UIA_IsDataValidForFormPropertyId                  UiaPropertyId = 30103
	UIA_ControllerForPropertyId                       UiaPropertyId = 30104
	UIA_DescribedByPropertyId                         UiaPropertyId = 30105
	UIA_FlowsToPropertyId                             UiaPropertyId = 30106
	UIA_ProviderDescriptionPropertyId                 UiaPropertyId = 30107
	UIA_IsItemContainerPatternAvailablePropertyId     UiaPropertyId = 30108
	UIA_IsVirtualizedItemPatternAvailablePropertyId   UiaPropertyId = 30109
	UIA_IsSynchronizedInputPatternAvailablePropertyId UiaPropertyId = 30110
	UIA_OptimizeForVisualContentPropertyId            UiaPropertyId = 30111
	UIA_IsObjectModelPatternAvailablePropertyId       UiaPropertyId = 30112
	UIA_AnnotationAnnotationTypeIdPropertyId          UiaPropertyId = 30113
	UIA_AnnotationAnnotationTypeNamePropertyId        UiaPropertyId = 30114
	UIA_AnnotationAuthorPropertyId                    UiaPropertyId = 30115
	UIA_AnnotationDateTimePropertyId                  UiaPropertyId = 30116
	UIA_AnnotationTargetPropertyId                    UiaPropertyId = 30117
	UIA_IsAnnotationPatternAvailablePropertyId        UiaPropertyId = 30118
	UIA_IsTextPattern2AvailablePropertyId             UiaPropertyId = 30119
	UIA_StylesStyleIdPropertyId                       UiaPropertyId = 30120
	UIA_StylesStyleNamePropertyId                     UiaPropertyId = 30121
	UIA_StylesFillColorPropertyId                     UiaPropertyId = 30122
	UIA_StylesFillPatternStylePropertyId              UiaPropertyId = 30123
	UIA_StylesShapePropertyId                         UiaPropertyId = 30124
	UIA_StylesFillPatternColorPropertyId              UiaPropertyId = 30125
	UIA_StylesExtendedPropertiesPropertyId            UiaPropertyId = 30126
	UIA_IsStylesPatternAvailablePropertyId            UiaPropertyId = 30127
	UIA_IsSpreadsheetPatternAvailablePropertyId       UiaPropertyId = 30128
	UIA_SpreadsheetItemFormulaPropertyId              UiaPropertyId = 30129
	UIA_SpreadsheetItemAnnotationObjectsPropertyId    UiaPropertyId = 30130
	UIA_SpreadsheetItemAnnotationTypesPropertyId      UiaPropertyId = 30131
	UIA_IsSpreadsheetItemPatternAvailablePropertyId   UiaPropertyId = 30132
	UIA_Transform2CanZoomPropertyId                   UiaPropertyId = 30133
	UIA_IsTransformPattern2AvailablePropertyId        UiaPropertyId = 30134
	UIA_LiveSettingPropertyId                         UiaPropertyId = 30135
	UIA_IsTextChildPatternAvailablePropertyId         UiaPropertyId = 30136
	UIA_IsDragPatternAvailablePropertyId              UiaPropertyId = 30137
	UIA_DragIsGrabbedPropertyId                       UiaPropertyId = 30138
	UIA_DragDropEffectPropertyId                      UiaPropertyId = 30139
	UIA_DragDropEffectsPropertyId                     UiaPropertyId = 30140
	UIA_IsDropTargetPatternAvailablePropertyId        UiaPropertyId = 30141
	UIA_DropTargetDropTargetEffectPropertyId          UiaPropertyId = 30142
	UIA_DropTargetDropTargetEffectsPropertyId         UiaPropertyId = 30143
	UIA_DragGrabbedItemsPropertyId                    UiaPropertyId = 30144
	UIA_Transform2ZoomLevelPropertyId                 UiaPropertyId = 30145
	UIA_Transform2ZoomMinimumPropertyId               UiaPropertyId = 30146
	UIA_Transform2ZoomMaximumPropertyId               UiaPropertyId = 30147
	UIA_FlowsFromPropertyId                           UiaPropertyId = 30148
	UIA_IsTextEditPatternAvailablePropertyId          UiaPropertyId = 30149
	UIA_IsPeripheralPropertyId                        UiaPropertyId = 30150
	UIA_IsCustomNavigationPatternAvailablePropertyId  UiaPropertyId = 30151
	UIA_PositionInSetPropertyId                       UiaPropertyId = 30152
	UIA_SizeOfSetPropertyId                           UiaPropertyId = 30153
	UIA_LevelPropertyId                               UiaPropertyId = 30154
	UIA_AnnotationTypesPropertyId                     UiaPropertyId = 30155
	UIA_AnnotationObjectsPropertyId                   UiaPropertyId = 30156
	UIA_LandmarkTypePropertyId                        UiaPropertyId = 30157
	UIA_LocalizedLandmarkTypePropertyId               UiaPropertyId = 30158
	UIA_FullDescriptionPropertyId                     UiaPropertyId = 30159
	UIA_FillColorPropertyId                           UiaPropertyId = 30160
	UIA_OutlineColorPropertyId                        UiaPropertyId = 30161
	UIA_FillTypePropertyId                            UiaPropertyId = 30162
	UIA_VisualEffectsPropertyId                       UiaPropertyId = 30163
	UIA_OutlineThicknessPropertyId                    UiaPropertyId = 30164
	UIA_CenterPointPropertyId                         UiaPropertyId = 30165
	UIA_RotationPropertyId                            UiaPropertyId = 30166
	UIA_SizePropertyId                                UiaPropertyId = 30167
	UIA_IsSelectionPattern2AvailablePropertyId        UiaPropertyId = 30168
	UIA_Selection2FirstSelectedItemPropertyId         UiaPropertyId = 30169
	UIA_Selection2LastSelectedItemPropertyId          UiaPropertyId = 30170
	UIA_Selection2CurrentSelectedItemPropertyId       UiaPropertyId = 30171
	UIA_Selection2ItemCountPropertyId                 UiaPropertyId = 30172
	UIA_HeadingLevelPropertyId                        UiaPropertyId = 30173
	UIA_IsDialogPropertyId                            UiaPropertyId = 30174
)

type UiaControlTypeId uint32

const (
	UIA_ButtonControlTypeId       UiaControlTypeId = 50000
	UIA_CalendarControlTypeId     UiaControlTypeId = 50001
	UIA_CheckBoxControlTypeId     UiaControlTypeId = 50002
	UIA_ComboBoxControlTypeId     UiaControlTypeId = 50003
	UIA_EditControlTypeId         UiaControlTypeId = 50004
	UIA_HyperlinkControlTypeId    UiaControlTypeId = 50005
	UIA_ImageControlTypeId        UiaControlTypeId = 50006
	UIA_ListItemControlTypeId     UiaControlTypeId = 50007
	UIA_ListControlTypeId         UiaControlTypeId = 50008
	UIA_MenuControlTypeId         UiaControlTypeId = 50009
	UIA_MenuBarControlTypeId      UiaControlTypeId = 50010
	UIA_MenuItemControlTypeId     UiaControlTypeId = 50011
	UIA_ProgressBarControlTypeId  UiaControlTypeId = 50012
	UIA_RadioButtonControlTypeId  UiaControlTypeId = 50013
	UIA_ScrollBarControlTypeId    UiaControlTypeId = 50014
	UIA_SliderControlTypeId       UiaControlTypeId = 50015
	UIA_SpinnerControlTypeId      UiaControlTypeId = 50016
	UIA_StatusBarControlTypeId    UiaControlTypeId = 50017
	UIA_TabControlTypeId          UiaControlTypeId = 50018
	UIA_TabItemControlTypeId      UiaControlTypeId = 50019
	UIA_TextControlTypeId         UiaControlTypeId = 50020
	UIA_ToolBarControlTypeId      UiaControlTypeId = 50021
	UIA_ToolTipControlTypeId      UiaControlTypeId = 50022
	UIA_TreeControlTypeId         UiaControlTypeId = 50023
	UIA_TreeItemControlTypeId     UiaControlTypeId = 50024
	UIA_CustomControlTypeId       UiaControlTypeId = 50025
	UIA_GroupControlTypeId        UiaControlTypeId = 50026
	UIA_ThumbControlTypeId        UiaControlTypeId = 50027
	UIA_DataGridControlTypeId     UiaControlTypeId = 50028
	UIA_DataItemControlTypeId     UiaControlTypeId = 50029
	UIA_DocumentControlTypeId     UiaControlTypeId = 50030
	UIA_SplitButtonControlTypeId  UiaControlTypeId = 50031
	UIA_WindowControlTypeId       UiaControlTypeId = 50032
	UIA_PaneControlTypeId         UiaControlTypeId = 50033
	UIA_HeaderControlTypeId       UiaControlTypeId = 50034
	UIA_HeaderItemControlTypeId   UiaControlTypeId = 50035
	UIA_TableControlTypeId        UiaControlTypeId = 50036
	UIA_TitleBarControlTypeId     UiaControlTypeId = 50037
	UIA_SeparatorControlTypeId    UiaControlTypeId = 50038
	UIA_SemanticZoomControlTypeId UiaControlTypeId = 50039
	UIA_AppBarControlTypeId       UiaControlTypeId = 50040
)

type UiaEventId uint32

const (
	UIA_ToolTipOpenedEventId                             UiaEventId = 20000
	UIA_ToolTipClosedEventId                             UiaEventId = 20001
	UIA_StructureChangedEventId                          UiaEventId = 20002
	UIA_MenuOpenedEventId                                UiaEventId = 20003
	UIA_AutomationPropertyChangedEventId                 UiaEventId = 20004
	UIA_AutomationFocusChangedEventId                    UiaEventId = 20005
	UIA_AsyncContentLoadedEventId                        UiaEventId = 20006
	UIA_MenuClosedEventId                                UiaEventId = 20007
	UIA_LayoutInvalidatedEventId                         UiaEventId = 20008
	UIA_Invoke_InvokedEventId                            UiaEventId = 20009
	UIA_SelectionItem_ElementAddedToSelectionEventId     UiaEventId = 20010
	UIA_SelectionItem_ElementRemovedFromSelectionEventId UiaEventId = 20011
	UIA_SelectionItem_ElementSelectedEventId             UiaEventId = 20012
	UIA_Selection_InvalidatedEventId                     UiaEventId = 20013
	UIA_Text_TextSelectionChangedEventId                 UiaEventId = 20014
	UIA_Text_TextChangedEventId                          UiaEventId = 20015
	UIA_Window_WindowOpenedEventId                       UiaEventId = 20016
	UIA_Window_WindowClosedEventId                       UiaEventId = 20017
	UIA_MenuModeStartEventId                             UiaEventId = 20018
	UIA_MenuModeEndEventId                               UiaEventId = 20019
	UIA_InputReachedTargetEventId                        UiaEventId = 20020
	UIA_InputReachedOtherElementEventId                  UiaEventId = 20021
	UIA_InputDiscardedEventId                            UiaEventId = 20022
	UIA_SystemAlertEventId                               UiaEventId = 20023
	UIA_LiveRegionChangedEventId                         UiaEventId = 20024
	UIA_HostedFragmentRootsInvalidatedEventId            UiaEventId = 20025
	UIA_Drag_DragStartEventId                            UiaEventId = 20026
	UIA_Drag_DragCancelEventId                           UiaEventId = 20027
	UIA_Drag_DragCompleteEventId                         UiaEventId = 20028
	UIA_DropTarget_DragEnterEventId                      UiaEventId = 20029
	UIA_DropTarget_DragLeaveEventId                      UiaEventId = 20030
	UIA_DropTarget_DroppedEventId                        UiaEventId = 20031
	UIA_TextEdit_TextChangedEventId                      UiaEventId = 20032
	UIA_TextEdit_ConversionTargetChangedEventId          UiaEventId = 20033
	UIA_ChangesEventId                                   UiaEventId = 20034
	UIA_NotificationEventId                              UiaEventId = 20035
	UIA_ActiveTextPositionChangedEventId                 UiaEventId = 20036
)

type StructureChangeType uint32

const (
	StructureChangeType_ChildAdded StructureChangeType = iota
	StructureChangeType_ChildRemoved
	StructureChangeType_ChildrenInvalidated
	StructureChangeType_ChildrenBulkAdded
	StructureChangeType_ChildrenBulkRemoved
	StructureChangeType_ChildrenReordered
)

type NavigateDirection uint32

const (
	NavigateDirection_Parent          NavigateDirection = 0
	NavigateDirection_NextSibling     NavigateDirection = 1
	NavigateDirection_PreviousSibling NavigateDirection = 2
	NavigateDirection_FirstChild      NavigateDirection = 3
	NavigateDirection_LastChild       NavigateDirection = 4
)

type NotificationKind uint32

const (
	NotificationKind_ItemAdded       NotificationKind = 0
	NotificationKind_ItemRemoved     NotificationKind = 1
	NotificationKind_ActionCompleted NotificationKind = 2
	NotificationKind_ActionAborted   NotificationKind = 3
	NotificationKind_Other           NotificationKind = 4
)

type NotificationProcessing uint32

const (
	NotificationProcessing_ImportantAll          NotificationProcessing = 0
	NotificationProcessing_ImportantMostRecent   NotificationProcessing = 1
	NotificationProcessing_All                   NotificationProcessing = 2
	NotificationProcessing_MostRecent            NotificationProcessing = 3
	NotificationProcessing_CurrentThenMostRecent NotificationProcessing = 4
)

type UiaRect struct {
	_ structs.HostLayout

	Left   float64
	Top    float64
	Width  float64
	Height float64
}

type UiaChangeInfo struct {
	_ structs.HostLayout

	uiaId     uint32
	payload   Variant
	extraInfo Variant
}

type SAFEARRAYBOUND struct {
	_ structs.HostLayout

	cElements uint64
	lLbound   int64
}

type SAFEARRAY struct {
	_ structs.HostLayout

	cDims      uint16
	fFeatures  SAFEARRAY_FLAGS
	cbElements uint64
	cLocks     uint64
	pvData     uintptr
	rgsabound  [1]SAFEARRAYBOUND
}

type SAFEARRAY_FLAGS uint16

const UiaAppendRuntimeId = 3

const (
	FADF_AUTO        SAFEARRAY_FLAGS = 0x0001
	FADF_STATIC      SAFEARRAY_FLAGS = 0x0002
	FADF_EMBEDDED    SAFEARRAY_FLAGS = 0x0004
	FADF_FIXEDSIZE   SAFEARRAY_FLAGS = 0x0010
	FADF_RECORD      SAFEARRAY_FLAGS = 0x0020
	FADF_HAVEIID     SAFEARRAY_FLAGS = 0x0040
	FADF_HAVEVARTYPE SAFEARRAY_FLAGS = 0x0080
	FADF_BSTR        SAFEARRAY_FLAGS = 0x0100
	FADF_UNKNOWN     SAFEARRAY_FLAGS = 0x0200
	FADF_DISPATCH    SAFEARRAY_FLAGS = 0x0400
	FADF_VARIANT     SAFEARRAY_FLAGS = 0x0800
	FADF_RESERVED    SAFEARRAY_FLAGS = 0xF008
)

type IUIAutomationTogglePatternVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Get_CachedToggleState  uintptr
	Get_CurrentToggleState uintptr
	Toggle                 uintptr
}

type IUIAutomationInvokePatternVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Invoke uintptr
}

type IRawElementProviderSimpleVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Get_ProviderOptions        uintptr
	GetPatternProvider         uintptr
	GetPropertyValue           uintptr
	Get_HostRawElementProvider uintptr
}

type IRawElementProviderFragmentVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Navigate                 uintptr
	GetRuntimeId             uintptr
	Get_BoundingRectangle    uintptr
	GetEmbeddedFragmentRoots uintptr
	SetFocus                 uintptr
	Get_FragmentRoot         uintptr
}

type IRawElementProviderFragmentRootVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	ElementProviderFromPoint uintptr
	GetFocus                 uintptr
}

type IInvokeProviderVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Invoke uintptr
}

type IValueProviderVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	SetValue       uintptr
	Get_Value      uintptr
	Get_IsReadOnly uintptr
}

type IScrollProviderVTable struct {
	_ structs.HostLayout
	IUnknownVTable

	Scroll                      uintptr
	SetScrollPercent            uintptr
	Get_HorizontalScrollPercent uintptr
	Get_VerticalScrollPercent   uintptr
	Get_HorizontalViewSize      uintptr
	Get_VerticalViewSize        uintptr
	Get_HorizontallyScrollable  uintptr
	Get_VerticallyScrollable    uintptr
}

type SemanticVTable struct {
	Simple       *IRawElementProviderSimpleVTable
	Fragment     *IRawElementProviderFragmentVTable
	FragmentRoot *IRawElementProviderFragmentRootVTable
	Invoke       *IInvokeProviderVTable
	Value        *IValueProviderVTable
	Scroll       *IScrollProviderVTable
}

type ProviderOptions uint32

const (
	ProviderOptions_ClientSideProvider     ProviderOptions = 0x1
	ProviderOptions_ServerSideProvider     ProviderOptions = 0x2
	ProviderOptions_NonClientAreaProvider  ProviderOptions = 0x4
	ProviderOptions_OverrideProvider       ProviderOptions = 0x8
	ProviderOptions_ProviderOwnsSetFocus   ProviderOptions = 0x10
	ProviderOptions_UseComThreading        ProviderOptions = 0x20
	ProviderOptions_RefuseNonClientSupport ProviderOptions = 0x40
	ProviderOptions_HasNativeIAccessible   ProviderOptions = 0x80
	ProviderOptions_UseClientCoordinates   ProviderOptions = 0x100

	CLSCTX_INPROC_SERVER     = 0x1
	COINIT_APARTMENTTHREADED = 0x2
	COINIT_DISABLE_OLE1DDE   = 0x4

	S_OK           = uintptr(0x00000000)
	E_ABORT        = uintptr(0x80004004)
	E_ACCESSDENIED = uintptr(0x80070005)
	E_FAIL         = uintptr(0x80004005)
	E_HANDLE       = uintptr(0x80070006)
	E_INVALIDARG   = uintptr(0x80070057)
	E_NOINTERFACE  = uintptr(0x80004002)
	E_NOTIMPL      = uintptr(0x80004001)
	E_OUTOFMEMORY  = uintptr(0x8007000E)
	E_POINTER      = uintptr(0x80004003)
	E_UNEXPECTED   = uintptr(0x8000FFFF)
)

const (
	TRUE = 1

	CPS_CANCEL = 0x0004

	CS_HREDRAW     = 0x0002
	CS_INSERTCHAR  = 0x2000
	CS_NOMOVECARET = 0x4000
	CS_VREDRAW     = 0x0001
	CS_OWNDC       = 0x0020

	CW_USEDEFAULT = -2147483648

	GWL_STYLE = ^(uintptr(16) - 1) // -16

	GCS_COMPSTR       = 0x0008
	GCS_COMPREADSTR   = 0x0001
	GCS_CURSORPOS     = 0x0080
	GCS_DELTASTART    = 0x0100
	GCS_RESULTREADSTR = 0x0200
	GCS_RESULTSTR     = 0x0800

	CFS_POINT        = 0x0002
	CFS_CANDIDATEPOS = 0x0040

	HWND_TOPMOST = ^(uint32(1) - 1) // -1

	HTCAPTION     = 2
	HTCLIENT      = 1
	HTLEFT        = 10
	HTRIGHT       = 11
	HTTOP         = 12
	HTTOPLEFT     = 13
	HTTOPRIGHT    = 14
	HTBOTTOM      = 15
	HTBOTTOMLEFT  = 16
	HTBOTTOMRIGHT = 17

	IDC_APPSTARTING = 32650 // Standard arrow and small hourglass
	IDC_ARROW       = 32512 // Standard arrow
	IDC_CROSS       = 32515 // Crosshair
	IDC_HAND        = 32649 // Hand
	IDC_HELP        = 32651 // Arrow and question mark
	IDC_IBEAM       = 32513 // I-beam
	IDC_NO          = 32648 // Slashed circle
	IDC_SIZEALL     = 32646 // Four-pointed arrow pointing north, south, east, and west
	IDC_SIZENESW    = 32643 // Double-pointed arrow pointing northeast and southwest
	IDC_SIZENS      = 32645 // Double-pointed arrow pointing north and south
	IDC_SIZENWSE    = 32642 // Double-pointed arrow pointing northwest and southeast
	IDC_SIZEWE      = 32644 // Double-pointed arrow pointing west and east
	IDC_UPARROW     = 32516 // Vertical arrow
	IDC_WAIT        = 32514 // Hour

	INFINITE = 0xFFFFFFFF

	LOGPIXELSX = 88

	MDT_EFFECTIVE_DPI = 0

	MONITOR_DEFAULTTOPRIMARY = 1

	NI_COMPOSITIONSTR = 0x0015

	SIZE_MAXIMIZED = 2
	SIZE_MINIMIZED = 1
	SIZE_RESTORED  = 0

	SCS_SETSTR = GCS_COMPREADSTR | GCS_COMPSTR

	SM_CXSIZEFRAME = 32
	SM_CYSIZEFRAME = 33

	SW_SHOWDEFAULT   = 10
	SW_SHOWMINIMIZED = 2
	SW_SHOWMAXIMIZED = 3
	SW_SHOWNORMAL    = 1
	SW_SHOW          = 5

	SWP_FRAMECHANGED  = 0x0020
	SWP_NOMOVE        = 0x0002
	SWP_NOOWNERZORDER = 0x0200
	SWP_NOSIZE        = 0x0001
	SWP_NOZORDER      = 0x0004
	SWP_SHOWWINDOW    = 0x0040

	USER_TIMER_MINIMUM = 0x0000000A

	VK_CONTROL = 0x11
	VK_LWIN    = 0x5B
	VK_MENU    = 0x12
	VK_RWIN    = 0x5C
	VK_SHIFT   = 0x10

	VK_BACK   = 0x08
	VK_DELETE = 0x2e
	VK_DOWN   = 0x28
	VK_END    = 0x23
	VK_ESCAPE = 0x1b
	VK_HOME   = 0x24
	VK_LEFT   = 0x25
	VK_NEXT   = 0x22
	VK_PRIOR  = 0x21
	VK_RIGHT  = 0x27
	VK_RETURN = 0x0d
	VK_SPACE  = 0x20
	VK_TAB    = 0x09
	VK_UP     = 0x26

	VK_F1  = 0x70
	VK_F2  = 0x71
	VK_F3  = 0x72
	VK_F4  = 0x73
	VK_F5  = 0x74
	VK_F6  = 0x75
	VK_F7  = 0x76
	VK_F8  = 0x77
	VK_F9  = 0x78
	VK_F10 = 0x79
	VK_F11 = 0x7A
	VK_F12 = 0x7B

	VK_OEM_1      = 0xba
	VK_OEM_PLUS   = 0xbb
	VK_OEM_COMMA  = 0xbc
	VK_OEM_MINUS  = 0xbd
	VK_OEM_PERIOD = 0xbe
	VK_OEM_2      = 0xbf
	VK_OEM_3      = 0xc0
	VK_OEM_4      = 0xdb
	VK_OEM_5      = 0xdc
	VK_OEM_6      = 0xdd
	VK_OEM_7      = 0xde
	VK_OEM_102    = 0xe2

	UNICODE_NOCHAR = 65535

	WM_CANCELMODE            = 0x001F
	WM_CHAR                  = 0x0102
	WM_CLOSE                 = 0x0010
	WM_COPYDATA              = 0x004A
	WM_CREATE                = 0x0001
	WM_DPICHANGED            = 0x02E0
	WM_DESTROY               = 0x0002
	WM_ERASEBKGND            = 0x0014
	WM_GETMINMAXINFO         = 0x0024
	WM_IME_COMPOSITION       = 0x010F
	WM_IME_ENDCOMPOSITION    = 0x010E
	WM_IME_STARTCOMPOSITION  = 0x010D
	WM_KEYDOWN               = 0x0100
	WM_KEYUP                 = 0x0101
	WM_KILLFOCUS             = 0x0008
	WM_LBUTTONDOWN           = 0x0201
	WM_LBUTTONUP             = 0x0202
	WM_MBUTTONDOWN           = 0x0207
	WM_MBUTTONUP             = 0x0208
	WM_MOUSEMOVE             = 0x0200
	WM_MOUSEWHEEL            = 0x020A
	WM_MOUSEHWHEEL           = 0x020E
	WM_NCACTIVATE            = 0x0086
	WM_NCHITTEST             = 0x0084
	WM_NCCALCSIZE            = 0x0083
	WM_PAINT                 = 0x000F
	WM_POINTERCAPTURECHANGED = 0x024C
	WM_POINTERDOWN           = 0x0246
	WM_POINTERUP             = 0x0247
	WM_POINTERUPDATE         = 0x0245
	WM_POINTERWHEEL          = 0x024E
	WM_POINTERHWHEEL         = 0x024F
	WM_QUIT                  = 0x0012
	WM_RBUTTONDOWN           = 0x0204
	WM_RBUTTONUP             = 0x0205
	WM_SETCURSOR             = 0x0020
	WM_SETFOCUS              = 0x0007
	WM_SHOWWINDOW            = 0x0018
	WM_SIZE                  = 0x0005
	WM_STYLECHANGED          = 0x007D
	WM_SYSKEYDOWN            = 0x0104
	WM_SYSKEYUP              = 0x0105
	WM_TIMER                 = 0x0113
	WM_UNICHAR               = 0x0109
	WM_USER                  = 0x0400
	WM_WINDOWPOSCHANGED      = 0x0047
	WM_GETOBJECT             = 0x003D

	WS_CLIPCHILDREN     = 0x02000000
	WS_CLIPSIBLINGS     = 0x04000000
	WS_MAXIMIZE         = 0x01000000
	WS_ICONIC           = 0x20000000
	WS_VISIBLE          = 0x10000000
	WS_OVERLAPPED       = 0x00000000
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME |
		WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_CAPTION     = 0x00C00000
	WS_SYSMENU     = 0x00080000
	WS_THICKFRAME  = 0x00040000
	WS_MINIMIZEBOX = 0x00020000
	WS_MAXIMIZEBOX = 0x00010000

	WS_EX_APPWINDOW  = 0x00040000
	WS_EX_WINDOWEDGE = 0x00000100

	QS_ALLINPUT = 0x04FF

	MWMO_WAITALL        = 0x0001
	MWMO_INPUTAVAILABLE = 0x0004

	WAIT_OBJECT_0 = 0

	PM_REMOVE   = 0x0001
	PM_NOREMOVE = 0x0000

	GHND = 0x0042

	CF_UNICODETEXT = 13
	IMAGE_BITMAP   = 0
	IMAGE_ICON     = 1
	IMAGE_CURSOR   = 2

	LR_CREATEDIBSECTION = 0x00002000
	LR_DEFAULTCOLOR     = 0x00000000
	LR_DEFAULTSIZE      = 0x00000040
	LR_LOADFROMFILE     = 0x00000010
	LR_LOADMAP3DCOLORS  = 0x00001000
	LR_LOADTRANSPARENT  = 0x00000020
	LR_MONOCHROME       = 0x00000001
	LR_SHARED           = 0x00008000
	LR_VGACOLOR         = 0x00000080
)

var (
	kernel32          = syscall.NewLazySystemDLL("kernel32.dll")
	_GetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
	_GlobalAlloc      = kernel32.NewProc("GlobalAlloc")
	_GlobalFree       = kernel32.NewProc("GlobalFree")
	_GlobalLock       = kernel32.NewProc("GlobalLock")
	_GlobalUnlock     = kernel32.NewProc("GlobalUnlock")

	user32                       = syscall.NewLazySystemDLL("user32.dll")
	_AdjustWindowRectEx          = user32.NewProc("AdjustWindowRectEx")
	_CallMsgFilter               = user32.NewProc("CallMsgFilterW")
	_CloseClipboard              = user32.NewProc("CloseClipboard")
	_CreateWindowEx              = user32.NewProc("CreateWindowExW")
	_DefWindowProc               = user32.NewProc("DefWindowProcW")
	_DestroyWindow               = user32.NewProc("DestroyWindow")
	_DispatchMessage             = user32.NewProc("DispatchMessageW")
	_FindWindow                  = user32.NewProc("FindWindowW")
	_EmptyClipboard              = user32.NewProc("EmptyClipboard")
	_EnableMouseInPointer        = user32.NewProc("EnableMouseInPointer")
	_GetWindowRect               = user32.NewProc("GetWindowRect")
	_GetClientRect               = user32.NewProc("GetClientRect")
	_GetClipboardData            = user32.NewProc("GetClipboardData")
	_GetDC                       = user32.NewProc("GetDC")
	_GetDpiForWindow             = user32.NewProc("GetDpiForWindow")
	_GetKeyState                 = user32.NewProc("GetKeyState")
	_GetMessage                  = user32.NewProc("GetMessageW")
	_GetMessageTime              = user32.NewProc("GetMessageTime")
	_GetMonitorInfo              = user32.NewProc("GetMonitorInfoW")
	_GetPointerInfo              = user32.NewProc("GetPointerInfo")
	_GetSystemMetrics            = user32.NewProc("GetSystemMetrics")
	_GetWindowLong               = user32.NewProc("GetWindowLongPtrW")
	_GetWindowLong32             = user32.NewProc("GetWindowLongW")
	_GetWindowPlacement          = user32.NewProc("GetWindowPlacement")
	_KillTimer                   = user32.NewProc("KillTimer")
	_LoadCursor                  = user32.NewProc("LoadCursorW")
	_LoadImage                   = user32.NewProc("LoadImageW")
	_MonitorFromPoint            = user32.NewProc("MonitorFromPoint")
	_MonitorFromWindow           = user32.NewProc("MonitorFromWindow")
	_MoveWindow                  = user32.NewProc("MoveWindow")
	_MsgWaitForMultipleObjectsEx = user32.NewProc("MsgWaitForMultipleObjectsEx")
	_OpenClipboard               = user32.NewProc("OpenClipboard")
	_PeekMessage                 = user32.NewProc("PeekMessageW")
	_PostMessage                 = user32.NewProc("PostMessageW")
	_PostQuitMessage             = user32.NewProc("PostQuitMessage")
	_ReleaseCapture              = user32.NewProc("ReleaseCapture")
	_RegisterClassExW            = user32.NewProc("RegisterClassExW")
	_RegisterTouchWindow         = user32.NewProc("RegisterTouchWindow")
	_ReleaseDC                   = user32.NewProc("ReleaseDC")
	_ScreenToClient              = user32.NewProc("ScreenToClient")
	_ShowWindow                  = user32.NewProc("ShowWindow")
	_SendMessage                 = user32.NewProc("SendMessageW")
	_SetCapture                  = user32.NewProc("SetCapture")
	_SetCursor                   = user32.NewProc("SetCursor")
	_SetClipboardData            = user32.NewProc("SetClipboardData")
	_SetForegroundWindow         = user32.NewProc("SetForegroundWindow")
	_SetFocus                    = user32.NewProc("SetFocus")
	_SetProcessDPIAware          = user32.NewProc("SetProcessDPIAware")
	_SetTimer                    = user32.NewProc("SetTimer")
	_SetWindowLong               = user32.NewProc("SetWindowLongPtrW")
	_SetWindowLong32             = user32.NewProc("SetWindowLongW")
	_SetWindowPlacement          = user32.NewProc("SetWindowPlacement")
	_SetWindowPos                = user32.NewProc("SetWindowPos")
	_SetWindowText               = user32.NewProc("SetWindowTextW")
	_TranslateMessage            = user32.NewProc("TranslateMessage")
	_UnregisterClass             = user32.NewProc("UnregisterClassW")
	_UpdateWindow                = user32.NewProc("UpdateWindow")

	shcore            = syscall.NewLazySystemDLL("shcore")
	_GetDpiForMonitor = shcore.NewProc("GetDpiForMonitor")

	gdi32          = syscall.NewLazySystemDLL("gdi32")
	_GetDeviceCaps = gdi32.NewProc("GetDeviceCaps")

	imm32                    = syscall.NewLazySystemDLL("imm32")
	_ImmGetContext           = imm32.NewProc("ImmGetContext")
	_ImmGetCompositionString = imm32.NewProc("ImmGetCompositionStringW")
	_ImmNotifyIME            = imm32.NewProc("ImmNotifyIME")
	_ImmReleaseContext       = imm32.NewProc("ImmReleaseContext")
	_ImmSetCandidateWindow   = imm32.NewProc("ImmSetCandidateWindow")
	_ImmSetCompositionWindow = imm32.NewProc("ImmSetCompositionWindow")

	dwmapi                        = syscall.NewLazySystemDLL("dwmapi")
	_DwmExtendFrameIntoClientArea = dwmapi.NewProc("DwmExtendFrameIntoClientArea")

	ole32             = syscall.NewLazySystemDLL("ole32")
	_CoCreateInstance = ole32.NewProc("CoCreateInstance")

	oleaut32               = syscall.NewLazySystemDLL("oleaut32")
	_VariantInit           = oleaut32.NewProc("VariantInit")
	_VariantClear          = oleaut32.NewProc("VariantClear")
	_VariantCopy           = oleaut32.NewProc("VariantCopy")
	_SafeArrayCreate       = oleaut32.NewProc("SafeArrayCreate")
	_SafeArrayCreateVector = oleaut32.NewProc("SafeArrayCreateVector")
	_SafeArrayPutElement   = oleaut32.NewProc("SafeArrayPutElement")
	_SafeArrayDestroy      = oleaut32.NewProc("SafeArrayDestroy")
	_SysAllocString        = oleaut32.NewProc("SysAllocString")

	uiautomationcore                        = syscall.NewLazySystemDLL("uiautomationcore")
	_UiaReturnRawElementProvider            = uiautomationcore.NewProc("UiaReturnRawElementProvider")
	_UiaRaiseChangesEvent                   = uiautomationcore.NewProc("UiaRaiseChangesEvent")
	_UiaRaiseNotificationEvent              = uiautomationcore.NewProc("UiaRaiseNotificationEvent")
	_UiaRaiseAutomationEvent                = uiautomationcore.NewProc("UiaRaiseAutomationEvent")
	_UiaRaiseAutomationPropertyChangedEvent = uiautomationcore.NewProc("UiaRaiseAutomationPropertyChangedEvent")
	_UiaRaiseStructureChangedEvent          = uiautomationcore.NewProc("UiaRaiseStructureChangedEvent")
	_UiaDisconnectProvider                  = uiautomationcore.NewProc("UiaDisconnectProvider")
	_UiaDisconnectAllProviders              = uiautomationcore.NewProc("UiaDisconnectAllProviders")
	_UiaHostProviderFromHwnd                = uiautomationcore.NewProc("UiaHostProviderFromHwnd")
)

func CoInitializeEx(pvReserved uintptr, dwCoInit uint32) error {
	err := syscall.CoInitializeEx(pvReserved, dwCoInit)
	if err != nil {
		return fmt.Errorf("CoInitializeEx failed: %w", err)
	}

	return nil
}

func CoUnintialize() {
	syscall.CoUninitialize()
}

func CoCreateInstance(rclsid *syscall.GUID, pUnkOuter uintptr, dwClsContext uint32, riid *syscall.GUID) (uintptr, error) {
	var ppv uintptr

	r, _, _ := _CoCreateInstance.Call(
		uintptr(unsafe.Pointer(rclsid)),
		pUnkOuter,
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppv)),
	)

	if r != 0 {
		return 0, fmt.Errorf("CoCreateInstance failed: %#x", r)
	}

	return ppv, nil
}

func SysAllocString(s string) (uintptr, error) {
	ptr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return 0, fmt.Errorf("converting string: %w", err)
	}

	ret, _, _ := _SysAllocString.Call(uintptr(unsafe.Pointer(ptr)))
	if ret == 0 {
		return 0, fmt.Errorf("SysAllocString failed: %#x", ret)
	}

	return ret, nil
}

func VariantInit() (Variant, error) {
	var pvarg Variant
	r, _, _ := _VariantInit.Call(uintptr(unsafe.Pointer(&pvarg)))

	if r != 0 {
		return Variant{}, fmt.Errorf("VariantInit failed: %#x", r)
	}

	return pvarg, nil
}

func VariantClear(pvarg *Variant) error {
	r, _, _ := _VariantClear.Call(uintptr(unsafe.Pointer(pvarg)))

	if r != 0 {
		return fmt.Errorf("VariantClear failed: %#x", r)
	}

	return nil
}

func VariantCopy(in *Variant) (Variant, error) {
	var out Variant
	r, _, _ := _VariantCopy.Call(
		uintptr(unsafe.Pointer(&out)),
		uintptr(unsafe.Pointer(in)),
	)

	if r != 0 {
		return Variant{}, fmt.Errorf("VariantCopy failed: %#x", r)
	}

	return out, nil
}

func SafeArrayCreate(vt VARIANT_TYPE, cDims uint32, rgsabound *SAFEARRAYBOUND) (*SAFEARRAY, error) {
	r, _, err := _SafeArrayCreate.Call(
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(rgsabound)),
	)

	sa, ok := any(r).(*SAFEARRAY)
	if !ok {
		return nil, fmt.Errorf("SafeArrayCreate failed: %v", err)
	}

	return sa, nil
}

func SafeArrayCreateVector(vt VARIANT_TYPE, lLbound int64, cElements uint64) (*SAFEARRAY, error) {
	r, _, err := _SafeArrayCreateVector.Call(
		uintptr(vt),
		uintptr(lLbound),
		uintptr(cElements),
	)

	sa, ok := any(r).(*SAFEARRAY)
	if !ok {
		return nil, fmt.Errorf("SafeArrayCreateVector failed: %v", err)
	}

	return sa, nil
}

func SafeArrayPutElement(sa *SAFEARRAY, rgIndices uint64, pv uintptr) error {
	r, _, _ := _SafeArrayPutElement.Call(
		uintptr(unsafe.Pointer(sa)),
		uintptr(unsafe.Pointer(&rgIndices)),
		uintptr(pv),
	)

	if r != S_OK {
		return fmt.Errorf("SafeArrayPutElement failed: %#x", r)
	}

	return nil
}

func SafeArrayDestroy(sa *SAFEARRAY) error {
	r, _, _ := _SafeArrayDestroy.Call(
		uintptr(unsafe.Pointer(sa)),
	)

	if r != S_OK {
		return fmt.Errorf("SafeArrayDestroy failed: %#x", r)
	}

	return nil
}

func UiaReturnRawElementProvider(hwnd syscall.Handle, wParam uintptr, lParam uintptr, el unsafe.Pointer) uintptr {
	w, _, _ := _UiaReturnRawElementProvider.Call(
		uintptr(hwnd),
		wParam,
		lParam,
		uintptr(el),
	)

	return w
}

func UiaRaiseChangesEvent(
	pProvider unsafe.Pointer,
	eventIdCount int32,
	pUiaChanges *UiaChangeInfo,
) error {
	r, _, _ := _UiaRaiseChangesEvent.Call(
		uintptr(unsafe.Pointer(pProvider)),
		uintptr(eventIdCount),
		uintptr(unsafe.Pointer(pUiaChanges)),
	)

	if r != 0 {
		return fmt.Errorf("UiaRaiseChangesEvent failed: %#x", r)
	}

	return nil
}

func UiaRaiseNotificationEvent(
	pProvider unsafe.Pointer,
	notificationKind NotificationKind,
	notificationProcessing NotificationProcessing,
	displayString *uint16,
	activityId *uint16,
) error {
	r, _, _ := _UiaRaiseNotificationEvent.Call(
		uintptr(unsafe.Pointer(pProvider)),
		uintptr(notificationKind),
		uintptr(notificationProcessing),
		uintptr(unsafe.Pointer(displayString)),
		uintptr(unsafe.Pointer(activityId)),
	)

	if r != 0 {
		return fmt.Errorf("UiaRaiseNotificationEvent failed: %#x", r)
	}

	return nil
}

func UiaRaiseAutomationEvent(pProvider unsafe.Pointer, id UiaEventId) uintptr {
	r, _, _ := _UiaRaiseAutomationEvent.Call(
		uintptr(unsafe.Pointer(pProvider)),
		uintptr(id),
	)

	if r != S_OK {
		return r
	}

	return S_OK
}

func UiaRaiseAutomationPropertyChangedEvent(
	pProvider unsafe.Pointer,
	id UiaPropertyId,
	oldV, newV *Variant,
) uintptr {
	r, _, _ := _UiaRaiseAutomationPropertyChangedEvent.Call(
		uintptr(pProvider),
		uintptr(id),
		uintptr(unsafe.Pointer(oldV)),
		uintptr(unsafe.Pointer(newV)),
	)

	if r != S_OK {
		return r
	}

	return S_OK
}

func UiaRaiseStructureChangedEvent(
	pProvider unsafe.Pointer,
	structureChangeType StructureChangeType,
	pRuntimeId unsafe.Pointer,
	cRuntimeIdLen uintptr,
) uintptr {
	r, _, _ := _UiaRaiseStructureChangedEvent.Call(
		uintptr(unsafe.Pointer(pProvider)),
		uintptr(structureChangeType),
		uintptr(pRuntimeId),
		uintptr(cRuntimeIdLen),
	)

	if r != S_OK {
		fmt.Println(r)
		return r
	}

	return S_OK

}

func UiaHostProviderFromHwnd(hwnd syscall.Handle, retVal *uintptr) uintptr {
	r, _, _ := _UiaHostProviderFromHwnd.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(retVal)),
	)

	if r != 0 {
		return r
	}

	return 0
}

func UiaDisconnectProvider(pProvider unsafe.Pointer) error {
	r, _, _ := _UiaDisconnectProvider.Call(uintptr(unsafe.Pointer(pProvider)))

	if r != 0 {
		return fmt.Errorf("UiaDisconnectProvider failed: %#x", r)
	}

	return nil
}

func UiaDisconnectAllProvider() error {
	r, _, _ := _UiaDisconnectAllProviders.Call()

	if r != 0 {
		return fmt.Errorf("UiaDisconnectAllProvider failed: %#x", r)
	}

	return nil
}

func AdjustWindowRectEx(r *Rect, dwStyle uint32, bMenu int, dwExStyle uint32) {
	_AdjustWindowRectEx.Call(uintptr(unsafe.Pointer(r)), uintptr(dwStyle), uintptr(bMenu), uintptr(dwExStyle))
}

func CallMsgFilter(m *Msg, nCode uintptr) bool {
	r, _, _ := _CallMsgFilter.Call(uintptr(unsafe.Pointer(m)), nCode)
	return r != 0
}

func CloseClipboard() error {
	r, _, err := _CloseClipboard.Call()
	if r == 0 {
		return fmt.Errorf("CloseClipboard: %v", err)
	}
	return nil
}

func CreateWindowEx(dwExStyle uint32, lpClassName uint16, lpWindowName string, dwStyle uint32, x, y, w, h int32, hWndParent, hMenu, hInstance syscall.Handle, lpParam uintptr) (syscall.Handle, error) {
	wname, err := syscall.UTF16PtrFromString(lpWindowName)
	if err != nil {
		return 0, fmt.Errorf("CreateWindowEx failed: %v", err)
	}
	hwnd, _, err := _CreateWindowEx.Call(
		uintptr(dwExStyle),
		uintptr(lpClassName),
		uintptr(unsafe.Pointer(wname)),
		uintptr(dwStyle),
		uintptr(x), uintptr(y),
		uintptr(w), uintptr(h),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))
	if hwnd == 0 {
		return 0, fmt.Errorf("CreateWindowEx failed: %v", err)
	}
	return syscall.Handle(hwnd), nil
}

func GetPointerInfo(pointerId uint32) (PointerInfo, error) {
	var info PointerInfo
	r1, _, err := _GetPointerInfo.Call(uintptr(pointerId), uintptr(unsafe.Pointer(&info)))
	if r1 == 0 {
		return PointerInfo{}, fmt.Errorf("GetPointerInfo failed: %v", err)
	}
	return info, nil
}

func RegisterTouchWindow(hwnd syscall.Handle, flags uint32) error {
	r1, _, err := _RegisterTouchWindow.Call(uintptr(hwnd), uintptr(flags))
	if r1 == 0 {
		return fmt.Errorf("RegisterTouchWindow failed: %v", err)
	}
	return nil
}

func EnableMouseInPointer(enable uint) error {
	r1, _, err := _EnableMouseInPointer.Call(uintptr(enable))
	if r1 == 0 {
		return fmt.Errorf("EnableMouseInPointer failed: %v", err)
	}
	return nil
}

func DefWindowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	r, _, _ := _DefWindowProc.Call(uintptr(hwnd), uintptr(msg), wparam, lparam)
	return r
}

func DestroyWindow(hwnd syscall.Handle) {
	_DestroyWindow.Call(uintptr(hwnd))
}

func DispatchMessage(m *Msg) {

	_DispatchMessage.Call(uintptr(unsafe.Pointer(m)))
}

func DwmExtendFrameIntoClientArea(hwnd syscall.Handle, margins Margins) error {
	r, _, _ := _DwmExtendFrameIntoClientArea.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&margins)))
	if r != 0 {
		return fmt.Errorf("DwmExtendFrameIntoClientArea: %#x", r)
	}
	return nil
}

func EmptyClipboard() error {
	r, _, err := _EmptyClipboard.Call()
	if r == 0 {
		return fmt.Errorf("EmptyClipboard: %v", err)
	}
	return nil
}

func FindWindow(lpClassName string) (syscall.Handle, error) {
	className, err := syscall.UTF16PtrFromString(lpClassName)
	if err != nil {
		return 0, fmt.Errorf("FindWindow failed: %v", err)
	}
	hwnd, _, err := _FindWindow.Call(uintptr(unsafe.Pointer(className)), 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("FindWindow failed: %v", err)
	}
	return syscall.Handle(hwnd), nil
}

func GetWindowRect(hwnd syscall.Handle) Rect {
	var r Rect
	_GetWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&r)))
	return r
}

func GetClientRect(hwnd syscall.Handle) Rect {
	var r Rect
	_GetClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&r)))
	return r
}

func GetClipboardData(format uint32) (syscall.Handle, error) {
	r, _, err := _GetClipboardData.Call(uintptr(format))
	if r == 0 {
		return 0, fmt.Errorf("GetClipboardData: %v", err)
	}
	return syscall.Handle(r), nil
}

func GetDC(hwnd syscall.Handle) (syscall.Handle, error) {
	hdc, _, err := _GetDC.Call(uintptr(hwnd))
	if hdc == 0 {
		return 0, fmt.Errorf("GetDC failed: %v", err)
	}
	return syscall.Handle(hdc), nil
}

func GetModuleHandle() (syscall.Handle, error) {
	h, _, err := _GetModuleHandleW.Call(uintptr(0))
	if h == 0 {
		return 0, fmt.Errorf("GetModuleHandleW failed: %v", err)
	}
	return syscall.Handle(h), nil
}

func getDeviceCaps(hdc syscall.Handle, index int32) int {
	c, _, _ := _GetDeviceCaps.Call(uintptr(hdc), uintptr(index))
	return int(c)
}

func getDpiForMonitor(hmonitor syscall.Handle, dpiType uint32) int {
	var dpiX, dpiY uintptr
	_GetDpiForMonitor.Call(uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
	return int(dpiX)
}

// GetSystemDPI returns the effective DPI of the system.
func GetSystemDPI() int {
	// Check for GetDpiForMonitor, introduced in Windows 8.1.
	if _GetDpiForMonitor.Find() == nil {
		hmon := monitorFromPoint(Point{}, MONITOR_DEFAULTTOPRIMARY)
		return getDpiForMonitor(hmon, MDT_EFFECTIVE_DPI)
	} else {
		// Fall back to the physical device DPI.
		screenDC, err := GetDC(0)
		if err != nil {
			return 96
		}
		defer ReleaseDC(screenDC)
		return getDeviceCaps(screenDC, LOGPIXELSX)
	}
}

func GetKeyState(nVirtKey int32) int16 {
	c, _, _ := _GetKeyState.Call(uintptr(nVirtKey))
	return int16(c)
}

func GetMessage(m *Msg, hwnd syscall.Handle, wMsgFilterMin, wMsgFilterMax uint32) int32 {
	r, _, _ := _GetMessage.Call(uintptr(unsafe.Pointer(m)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax))
	return int32(r)
}

func GetMessageTime() time.Duration {
	r, _, _ := _GetMessageTime.Call()
	return time.Duration(r) * time.Millisecond
}

func GetSystemMetrics(nIndex int) int {
	r, _, _ := _GetSystemMetrics.Call(uintptr(nIndex))
	return int(r)
}

// GetWindowDPI returns the effective DPI of the window.
func GetWindowDPI(hwnd syscall.Handle) int {
	// Check for GetDpiForWindow, introduced in Windows 10.
	if _GetDpiForWindow.Find() == nil {
		dpi, _, _ := _GetDpiForWindow.Call(uintptr(hwnd))
		return int(dpi)
	} else {
		return GetSystemDPI()
	}
}

func GetWindowPlacement(hwnd syscall.Handle) *WindowPlacement {
	var wp WindowPlacement
	wp.length = uint32(unsafe.Sizeof(wp))
	_GetWindowPlacement.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&wp)))
	return &wp
}

func GetMonitorInfo(hwnd syscall.Handle) MonitorInfo {
	var mi MonitorInfo
	mi.cbSize = uint32(unsafe.Sizeof(mi))
	v, _, _ := _MonitorFromWindow.Call(uintptr(hwnd), MONITOR_DEFAULTTOPRIMARY)
	_GetMonitorInfo.Call(v, uintptr(unsafe.Pointer(&mi)))
	return mi
}

func GetWindowLong(hwnd syscall.Handle, index uintptr) (val uintptr) {
	if runtime.GOARCH == "386" {
		val, _, _ = _GetWindowLong32.Call(uintptr(hwnd), index)
	} else {
		val, _, _ = _GetWindowLong.Call(uintptr(hwnd), index)
	}
	return
}

func ImmGetContext(hwnd syscall.Handle) syscall.Handle {
	h, _, _ := _ImmGetContext.Call(uintptr(hwnd))
	return syscall.Handle(h)
}

func ImmReleaseContext(hwnd, imc syscall.Handle) {
	_ImmReleaseContext.Call(uintptr(hwnd), uintptr(imc))
}

func ImmNotifyIME(imc syscall.Handle, action, index, value int) {
	_ImmNotifyIME.Call(uintptr(imc), uintptr(action), uintptr(index), uintptr(value))
}

func ImmGetCompositionString(imc syscall.Handle, key int) string {
	size, _, _ := _ImmGetCompositionString.Call(uintptr(imc), uintptr(key), 0, 0)
	if int32(size) <= 0 {
		return ""
	}
	u16 := make([]uint16, size/unsafe.Sizeof(uint16(0)))
	_ImmGetCompositionString.Call(uintptr(imc), uintptr(key), uintptr(unsafe.Pointer(&u16[0])), size)
	return string(utf16.Decode(u16))
}

func ImmGetCompositionValue(imc syscall.Handle, key int) int {
	val, _, _ := _ImmGetCompositionString.Call(uintptr(imc), uintptr(key), 0, 0)
	return int(int32(val))
}

func ImmSetCompositionWindow(imc syscall.Handle, x, y int) {
	f := CompositionForm{
		dwStyle: CFS_POINT,
		ptCurrentPos: Point{
			X: int32(x), Y: int32(y),
		},
	}
	_ImmSetCompositionWindow.Call(uintptr(imc), uintptr(unsafe.Pointer(&f)))
}

func ImmSetCandidateWindow(imc syscall.Handle, x, y int) {
	f := CandidateForm{
		dwStyle: CFS_CANDIDATEPOS,
		ptCurrentPos: Point{
			X: int32(x), Y: int32(y),
		},
	}
	_ImmSetCandidateWindow.Call(uintptr(imc), uintptr(unsafe.Pointer(&f)))
}

func SetWindowLong(hwnd syscall.Handle, idx uintptr, style uintptr) {
	if runtime.GOARCH == "386" {
		_SetWindowLong32.Call(uintptr(hwnd), idx, style)
	} else {
		_SetWindowLong.Call(uintptr(hwnd), idx, style)
	}
}

func SetWindowPlacement(hwnd syscall.Handle, wp *WindowPlacement) {
	_SetWindowPlacement.Call(uintptr(hwnd), uintptr(unsafe.Pointer(wp)))
}

func SetWindowPos(hwnd syscall.Handle, hwndInsertAfter uint32, x, y, dx, dy int32, style uintptr) {
	_SetWindowPos.Call(uintptr(hwnd), uintptr(hwndInsertAfter),
		uintptr(x), uintptr(y),
		uintptr(dx), uintptr(dy),
		style,
	)
}

func SetWindowText(hwnd syscall.Handle, title string) {
	wname, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		panic(err)
	}
	_SetWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(wname)))
}

func GlobalAlloc(size int) (syscall.Handle, error) {
	r, _, err := _GlobalAlloc.Call(GHND, uintptr(size))
	if r == 0 {
		return 0, fmt.Errorf("GlobalAlloc: %v", err)
	}
	return syscall.Handle(r), nil
}

func GlobalFree(h syscall.Handle) {
	_GlobalFree.Call(uintptr(h))
}

func GlobalLock(h syscall.Handle) (unsafe.Pointer, error) {
	r, _, err := _GlobalLock.Call(uintptr(h))
	if r == 0 {
		return nil, fmt.Errorf("GlobalLock: %v", err)
	}
	return unsafe.Pointer(r), nil
}

func GlobalUnlock(h syscall.Handle) {
	_GlobalUnlock.Call(uintptr(h))
}

func KillTimer(hwnd syscall.Handle, nIDEvent uintptr) error {
	r, _, err := _SetTimer.Call(uintptr(hwnd), uintptr(nIDEvent), 0, 0)
	if r == 0 {
		return fmt.Errorf("KillTimer failed: %v", err)
	}
	return nil
}

func LoadCursor(curID uint16) (syscall.Handle, error) {
	h, _, err := _LoadCursor.Call(0, uintptr(curID))
	if h == 0 {
		return 0, fmt.Errorf("LoadCursorW failed: %v", err)
	}
	return syscall.Handle(h), nil
}

func LoadImage(hInst syscall.Handle, res uint32, typ uint32, cx, cy int, fuload uint32) (syscall.Handle, error) {
	h, _, err := _LoadImage.Call(uintptr(hInst), uintptr(res), uintptr(typ), uintptr(cx), uintptr(cy), uintptr(fuload))
	if h == 0 {
		return 0, fmt.Errorf("LoadImageW failed: %v", err)
	}
	return syscall.Handle(h), nil
}

func MoveWindow(hwnd syscall.Handle, x, y, width, height int32, repaint bool) {
	var paint uintptr
	if repaint {
		paint = TRUE
	}
	_MoveWindow.Call(uintptr(hwnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height), paint)
}

func monitorFromPoint(pt Point, flags uint32) syscall.Handle {
	r, _, _ := _MonitorFromPoint.Call(uintptr(pt.X), uintptr(pt.Y), uintptr(flags))
	return syscall.Handle(r)
}

func MsgWaitForMultipleObjectsEx(nCount uint32, pHandles uintptr, millis, mask, flags uint32) (uint32, error) {
	r, _, err := _MsgWaitForMultipleObjectsEx.Call(uintptr(nCount), pHandles, uintptr(millis), uintptr(mask), uintptr(flags))
	res := uint32(r)
	if res == 0xFFFFFFFF {
		return 0, fmt.Errorf("MsgWaitForMultipleObjectsEx failed: %v", err)
	}
	return res, nil
}

func OpenClipboard(hwnd syscall.Handle) error {
	r, _, err := _OpenClipboard.Call(uintptr(hwnd))
	if r == 0 {
		return fmt.Errorf("OpenClipboard: %v", err)
	}
	return nil
}

func PeekMessage(m *Msg, hwnd syscall.Handle, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	r, _, _ := _PeekMessage.Call(uintptr(unsafe.Pointer(m)), uintptr(hwnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax), uintptr(wRemoveMsg))
	return r != 0
}

func PostQuitMessage(exitCode uintptr) {
	_PostQuitMessage.Call(exitCode)
}

func PostMessage(hwnd syscall.Handle, msg uint32, wParam, lParam uintptr) error {
	r, _, err := _PostMessage.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	if r == 0 {
		return fmt.Errorf("PostMessage failed: %v", err)
	}
	return nil
}

func ReleaseCapture() bool {
	r, _, _ := _ReleaseCapture.Call()
	return r != 0
}

func RegisterClassEx(cls *WndClassEx) (uint16, error) {
	a, _, err := _RegisterClassExW.Call(uintptr(unsafe.Pointer(cls)))
	if a == 0 {
		return 0, fmt.Errorf("RegisterClassExW failed: %v", err)
	}
	return uint16(a), nil
}

func ReleaseDC(hdc syscall.Handle) {
	_ReleaseDC.Call(uintptr(hdc))
}

func SendMessage(hwnd syscall.Handle, msg uint32, wParam, lParam uintptr) error {
	r, _, err := _SendMessage.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	if r == 0 {
		return fmt.Errorf("SendMessage failed: %v", err)
	}
	return nil
}

func SetForegroundWindow(hwnd syscall.Handle) {
	_SetForegroundWindow.Call(uintptr(hwnd))
}

func SetFocus(hwnd syscall.Handle) {
	_SetFocus.Call(uintptr(hwnd))
}

func SetProcessDPIAware() {
	_SetProcessDPIAware.Call()
}

func SetCapture(hwnd syscall.Handle) syscall.Handle {
	r, _, _ := _SetCapture.Call(uintptr(hwnd))
	return syscall.Handle(r)
}

func SetClipboardData(format uint32, mem syscall.Handle) error {
	r, _, err := _SetClipboardData.Call(uintptr(format), uintptr(mem))
	if r == 0 {
		return fmt.Errorf("SetClipboardData: %v", err)
	}
	return nil
}

func SetCursor(h syscall.Handle) {
	_SetCursor.Call(uintptr(h))
}

func SetTimer(hwnd syscall.Handle, nIDEvent uintptr, uElapse uint32, timerProc uintptr) error {
	r, _, err := _SetTimer.Call(uintptr(hwnd), uintptr(nIDEvent), uintptr(uElapse), timerProc)
	if r == 0 {
		return fmt.Errorf("SetTimer failed: %v", err)
	}
	return nil
}

func ScreenToClient(hwnd syscall.Handle, p *Point) {
	_ScreenToClient.Call(uintptr(hwnd), uintptr(unsafe.Pointer(p)))
}

func ShowWindow(hwnd syscall.Handle, nCmdShow int32) {
	_ShowWindow.Call(uintptr(hwnd), uintptr(nCmdShow))
}

func TranslateMessage(m *Msg) {
	_TranslateMessage.Call(uintptr(unsafe.Pointer(m)))
}

func UnregisterClass(cls uint16, hInst syscall.Handle) {
	_UnregisterClass.Call(uintptr(cls), uintptr(hInst))
}

func UpdateWindow(hwnd syscall.Handle) {
	_UpdateWindow.Call(uintptr(hwnd))
}

func (p WindowPlacement) Rect() Rect {
	return p.rcNormalPosition
}

func (p WindowPlacement) IsMinimized() bool {
	return p.showCmd == SW_SHOWMINIMIZED
}

func (p WindowPlacement) IsMaximized() bool {
	return p.showCmd == SW_SHOWMAXIMIZED
}

func (p *WindowPlacement) Set(Left, Top, Right, Bottom int) {
	p.rcNormalPosition.Left = int32(Left)
	p.rcNormalPosition.Top = int32(Top)
	p.rcNormalPosition.Right = int32(Right)
	p.rcNormalPosition.Bottom = int32(Bottom)
}

func (v *Variant) SetBool(b bool) {
	v.VT = VT_BOOL
	if b {
		v.Val = ^uintptr(0)
	} else {
		v.Val = 0
	}
}

func (v *Variant) SetInt32(i32 int32) {
	v.VT = VT_I4
	v.Val = uintptr(i32)
}

func (v *Variant) SetString(str string) {
	v.VT = VT_BSTR
	bstr, err := SysAllocString(str)
	if err != nil {
		panic("allocating string failed")
	}

	v.Val = bstr
}

func (v *Variant) ToGoValue() any {
	switch v.VT {
	case VT_I4:
		return int32(v.Val)
	case VT_BSTR:
		return syscall.UTF16PtrToString((*uint16)(unsafe.Pointer(&v.Val)))
	case VT_BOOL:
		return v.Val != 0
	case VT_UNKNOWN:
		return unsafe.Pointer(&v.Val)
	default:
		return nil
	}
}
