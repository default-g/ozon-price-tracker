package ozon

const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
const DefaultWindowWidth = 1920
const DefaultWindowHeight = 1080
const DefaultDisableGpuFlag = false
const DefaultHeadlessFlag = true
const DefaultEnableAutomationFlag = true
const DefaultDisableBlinkFeatures = "AutomationControlled"

type OzonClientBuilder struct {
    userAgent            string
    windowWidth          int
    windowHeight         int
    disableGpuFlag       bool
    headlessFlag         bool
    enableAutomationFlag bool
    disableBlinkFeatures string
}

func NewOzonClientBuilder() *OzonClientBuilder {
    return &OzonClientBuilder{
        userAgent:            DefaultUserAgent,
        windowWidth:          DefaultWindowWidth,
        windowHeight:         DefaultWindowHeight,
        disableGpuFlag:       DefaultDisableGpuFlag,
        headlessFlag:         DefaultHeadlessFlag,
        enableAutomationFlag: DefaultEnableAutomationFlag,
        disableBlinkFeatures: DefaultDisableBlinkFeatures,
    }
}

func (o *OzonClientBuilder) WithUserAgent(userAgent string) *OzonClientBuilder {
    o.userAgent = userAgent
    return o
}

func (o *OzonClientBuilder) WithWindowWidth(windowWidth int) *OzonClientBuilder {
    o.windowWidth = windowWidth
    return o
}

func (o *OzonClientBuilder) WithWindowHeight(windowHeight int) *OzonClientBuilder {
    o.windowHeight = windowHeight
    return o
}

func (o *OzonClientBuilder) WithDisableGpuFlag(disableGpuFlag bool) *OzonClientBuilder {
    o.disableGpuFlag = disableGpuFlag
    return o
}

func (o *OzonClientBuilder) WithHeadlessFlag(headlessFlag bool) *OzonClientBuilder {
    o.headlessFlag = headlessFlag
    return o
}

func (o *OzonClientBuilder) WithEnableAutomationFlag(enableAutomationFlag bool) *OzonClientBuilder {
    o.enableAutomationFlag = enableAutomationFlag
    return o
}

func (o *OzonClientBuilder) WithDisableBlinkFeatures(disableBlinkFeatures string) *OzonClientBuilder {
    o.disableBlinkFeatures = disableBlinkFeatures
    return o
}

func (o *OzonClientBuilder) Build() *ClientImpl {
    return NewClient(
        o.disableGpuFlag,
        o.enableAutomationFlag,
        o.disableBlinkFeatures,
        o.headlessFlag,
        o.windowWidth,
        o.windowHeight,
        o.userAgent,
    )
}
