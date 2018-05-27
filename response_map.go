package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

var respMap = map[string]response{
	"GetVersion":                 GetVersionResponse{},
	"GetAuthRequired":            GetAuthRequiredResponse{},
	"Authenticate":               AuthenticateResponse{},
	"SetHeartbeat":               SetHeartbeatResponse{},
	"SetFilenameFormatting":      SetFilenameFormattingResponse{},
	"GetFilenameFormatting":      GetFilenameFormattingResponse{},
	"SetCurrentProfile":          SetCurrentProfileResponse{},
	"GetCurrentProfile":          GetCurrentProfileResponse{},
	"ListProfiles":               ListProfilesResponse{},
	"StartStopRecording":         StartStopRecordingResponse{},
	"StartRecording":             StartRecordingResponse{},
	"StopRecording":              StopRecordingResponse{},
	"SetRecordingFolder":         SetRecordingFolderResponse{},
	"GetRecordingFolder":         GetRecordingFolderResponse{},
	"StartStopReplayBuffer":      StartStopReplayBufferResponse{},
	"StartReplayBuffer":          StartReplayBufferResponse{},
	"StopReplayBuffer":           StopReplayBufferResponse{},
	"SaveReplayBuffer":           SaveReplayBufferResponse{},
	"SetCurrentSceneCollection":  SetCurrentSceneCollectionResponse{},
	"GetCurrentSceneCollection":  GetCurrentSceneCollectionResponse{},
	"ListSceneCollections":       ListSceneCollectionsResponse{},
	"GetSceneItemProperties":     GetSceneItemPropertiesResponse{},
	"SetSceneItemProperties":     SetSceneItemPropertiesResponse{},
	"ResetSceneItem":             ResetSceneItemResponse{},
	"SetSceneItemRender":         SetSceneItemRenderResponse{},
	"SetSceneItemPosition":       SetSceneItemPositionResponse{},
	"SetSceneItemTransform":      SetSceneItemTransformResponse{},
	"SetSceneItemCrop":           SetSceneItemCropResponse{},
	"SetCurrentScene":            SetCurrentSceneResponse{},
	"GetCurrentScene":            GetCurrentSceneResponse{},
	"GetSceneList":               GetSceneListResponse{},
	"SetSceneItemOrder":          SetSceneItemOrderResponse{},
	"GetSourcesList":             GetSourcesListResponse{},
	"GetSourcesTypesList":        GetSourcesTypesListResponse{},
	"GetVolume":                  GetVolumeResponse{},
	"SetVolume":                  SetVolumeResponse{},
	"GetMute":                    GetMuteResponse{},
	"SetMute":                    SetMuteResponse{},
	"ToggleMute":                 ToggleMuteResponse{},
	"SetSyncOffset":              SetSyncOffsetResponse{},
	"GetSyncOffset":              GetSyncOffsetResponse{},
	"GetSourceSettings":          GetSourceSettingsResponse{},
	"SetSourceSettings":          SetSourceSettingsResponse{},
	"GetTextGDIPlusProperties":   GetTextGDIPlusPropertiesResponse{},
	"SetTextGDIPlusProperties":   SetTextGDIPlusPropertiesResponse{},
	"GetTextFreetype2Properties": GetTextFreetype2PropertiesResponse{},
	"SetTextFreetype2Properties": SetTextFreetype2PropertiesResponse{},
	"GetBrowserSourceProperties": GetBrowserSourcePropertiesResponse{},
	"SetBrowserSourceProperties": SetBrowserSourcePropertiesResponse{},
	"DeleteSceneItem":            DeleteSceneItemResponse{},
	"DuplicateSceneItem":         DuplicateSceneItemResponse{},
	"GetSpecialSources":          GetSpecialSourcesResponse{},
	"GetStreamingStatus":         GetStreamingStatusResponse{},
	"StartStopStreaming":         StartStopStreamingResponse{},
	"StartStreaming":             StartStreamingResponse{},
	"StopStreaming":              StopStreamingResponse{},
	"SetStreamSettings":          SetStreamSettingsResponse{},
	"GetStreamSettings":          GetStreamSettingsResponse{},
	"SaveStreamSettings":         SaveStreamSettingsResponse{},
	"GetStudioModeStatus":        GetStudioModeStatusResponse{},
	"GetPreviewScene":            GetPreviewSceneResponse{},
	"SetPreviewScene":            SetPreviewSceneResponse{},
	"TransitionToProgram":        TransitionToProgramResponse{},
	"EnableStudioMode":           EnableStudioModeResponse{},
	"DisableStudioMode":          DisableStudioModeResponse{},
	"ToggleStudioMode":           ToggleStudioModeResponse{},
	"GetTransitionList":          GetTransitionListResponse{},
	"GetCurrentTransition":       GetCurrentTransitionResponse{},
	"SetCurrentTransition":       SetCurrentTransitionResponse{},
	"SetTransitionDuration":      SetTransitionDurationResponse{},
	"GetTransitionDuration":      GetTransitionDurationResponse{},
}
