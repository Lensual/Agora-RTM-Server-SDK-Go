package agorartm

/*
#include <sys/mman.h>
#include <unistd.h>
#include <errno.h>
#include <signal.h>
#include <setjmp.h>
#include <string.h>
#include <stdint.h>

int is_valid_memory(const void* ptr) {
    if (ptr == NULL) return 0;

    // 获取页面大小
    long page_size = sysconf(_SC_PAGESIZE);
    if (page_size <= 0) return 0;

    // 计算指针所在页面的起始地址
    void* page_start = (void*)((uintptr_t)ptr & ~(page_size - 1));

    // 检查页面是否可访问
    if (msync(page_start, page_size, MS_ASYNC) == -1) {
        if (errno == ENOMEM) {
            return 0; // 页面不存在或不可访问
        }
    }

    return 1;
}

static jmp_buf segv_buf;

void segv_handler(int sig) {
    longjmp(segv_buf, 1);
}

int safe_strlen(const char* str) {
    struct sigaction old_action, new_action;

    // 设置信号处理器
    new_action.sa_handler = segv_handler;
    sigemptyset(&new_action.sa_mask);
    new_action.sa_flags = 0;

    sigaction(SIGSEGV, &new_action, &old_action);

    int result = 0;
    if (setjmp(segv_buf) == 0) {
        result = strlen(str);
    } else {
        result = -1; // 段错误
    }

    // 恢复原来的信号处理器
    sigaction(SIGSEGV, &old_action, NULL);

    return result;
}
*/
import "C"
import "unsafe"

func IsValidMemory(ptr unsafe.Pointer) bool {
	return C.is_valid_memory(ptr) != 0
}

// FastSafeCGoString - 快速版本，只做基本检查
func FastSafeCGoString(cstr *C.char) string {
	if cstr == nil {
		return ""
	}

	// 只做简单的内存页面检查，避免信号处理开销
	if C.is_valid_memory(unsafe.Pointer(cstr)) == 0 {
		return ""
	}

	return C.GoString(cstr)
}

// SafeCGoString - 完整安全检查版本
func SafeCGoString(cstr *C.char) string {
	if cstr == nil {
		return ""
	}

	// 检查内存是否可访问
	if C.is_valid_memory(unsafe.Pointer(cstr)) == 0 {
		return ""
	}

	length := C.safe_strlen(cstr)
	if length < 0 {
		return ""
	}

	return C.GoString(cstr)
}

// C.struct_C_UserList to UserList
func CUserListToUserList(cUserList *C.struct_C_UserList) *UserList {
	if cUserList == nil || cUserList.users == nil || cUserList.userCount == 0 {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cUserList.users)) {
		return nil
	}

	userCount := int(cUserList.userCount)
	users := make([]string, userCount)

	// 使用 unsafe.Slice 创建字符串指针切片
	cUsers := unsafe.Slice((**C.char)(unsafe.Pointer(cUserList.users)), userCount)

	for i := 0; i < userCount; i++ {
		if cUsers[i] != nil {
			users[i] = FastSafeCGoString(cUsers[i])
		} else {
			users[i] = ""
		}
	}

	return &UserList{
		Users:     users,
		UserCount: uint(userCount),
	}
}

// C.struct_C_Metadata to IMetadata
func CMetadataToIMetadata(cMetadata *C.struct_C_Metadata) *IMetadata {
	if cMetadata == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cMetadata)) {
		return nil
	}

	itemCount := uint32(cMetadata.itemCount)

	var itemsPtr unsafe.Pointer
	if cMetadata.items != nil {
		itemsPtr = unsafe.Pointer(cMetadata.items)
	}

	majorRevision := int64(cMetadata.majorRevision)

	if itemsPtr == nil || itemCount == 0 || !IsValidMemory(itemsPtr) {
		return &IMetadata{
			majorRevision: majorRevision,
			items:         make([]MetadataItem, 0),
			itemCount:     uint(itemCount),
		}
	}

	items := make([]MetadataItem, itemCount)

	// 使用 unsafe.Slice 创建 MetadataItem 指针切片
	cItems := unsafe.Slice((**C.struct_C_MetadataItem)(itemsPtr), itemCount)

	for i := 0; i < int(itemCount); i++ {
		if cItems[i] != nil && IsValidMemory(unsafe.Pointer(cItems[i])) {
			items[i] = MetadataItem{
				Key:          FastSafeCGoString(cItems[i].key),
				Value:        FastSafeCGoString(cItems[i].value),
				AuthorUserId: FastSafeCGoString(cItems[i].authorUserId),
				Revision:     uint64(cItems[i].revision),
				UpdateTs:     uint64(cItems[i].updateTs),
			}
		}
	}

	return &IMetadata{
		majorRevision: majorRevision,
		items:         items,
		itemCount:     uint(itemCount),
	}
}

// C.struct_C_LockDetail to LockDetail
func CLockDetailToLockDetail(cLockDetail *C.struct_C_LockDetail) *LockDetail {
	if cLockDetail == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cLockDetail)) {
		return nil
	}

	return &LockDetail{
		lockName: FastSafeCGoString(cLockDetail.lockName),
		owner:    FastSafeCGoString(cLockDetail.owner),
		ttl:      uint32(cLockDetail.ttl),
	}
}

// C.struct_C_UserState to UserState
func CUserStateToUserState(cUserState *C.struct_C_UserState) *UserState {
	if cUserState == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cUserState)) {
		return nil
	}

	states := make([]StateItem, cUserState.statesCount)

	if cUserState.states == nil || cUserState.statesCount == 0 || !IsValidMemory(unsafe.Pointer(cUserState.states)) {
		return &UserState{
			UserId:      C.GoString(cUserState.userId),
			States:      make([]StateItem, 0),
			StatesCount: 0,
		}
	}

	// 使用 unsafe.Slice 创建 StateItem 指针切片
	cStates := unsafe.Slice((**C.struct_C_StateItem)(unsafe.Pointer(cUserState.states)), cUserState.statesCount)

	for i := 0; i < int(cUserState.statesCount); i++ {
		if cStates[i] != nil && IsValidMemory(unsafe.Pointer(cStates[i])) {
			states[i] = StateItem{
				Key:   FastSafeCGoString(cStates[i].key),
				Value: FastSafeCGoString(cStates[i].value),
			}
		}
	}

	return &UserState{
		UserId:      FastSafeCGoString(cUserState.userId),
		States:      states,
		StatesCount: uint(cUserState.statesCount),
	}
}

// C.struct_C_ChannelInfo to ChannelInfo
func CChannelInfoToChannelInfo(cChannelInfo *C.struct_C_ChannelInfo) *ChannelInfo {
	if cChannelInfo == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cChannelInfo)) {
		return nil
	}

	return &ChannelInfo{
		ChannelName: FastSafeCGoString(cChannelInfo.channelName),
		ChannelType: RTM_CHANNEL_TYPE(cChannelInfo.channelType),
	}
}

// C.struct_C_LinkStateEvent to LinkStateEvent
func CLinkStateEventToLinkStateEvent(cLinkStateEvent *C.struct_C_LinkStateEvent) *LinkStateEvent {
	if cLinkStateEvent == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cLinkStateEvent)) {
		return nil
	}

	affectedChannels := make([]string, cLinkStateEvent.affectedChannelCount)

	// 使用 unsafe.Slice 创建字符串指针切片
	cAffectedChannels := unsafe.Slice((**C.char)(unsafe.Pointer(cLinkStateEvent.affectedChannels)), cLinkStateEvent.affectedChannelCount)

	for i := 0; i < int(cLinkStateEvent.affectedChannelCount); i++ {
		if cAffectedChannels[i] != nil {
			affectedChannels[i] = FastSafeCGoString(cAffectedChannels[i])
		} else {
			affectedChannels[i] = ""
		}
	}

	unrestoredChannels := make([]string, cLinkStateEvent.unrestoredChannelCount)

	// 使用 unsafe.Slice 创建字符串指针切片
	cUnrestoredChannels := unsafe.Slice((**C.char)(unsafe.Pointer(cLinkStateEvent.unrestoredChannels)), cLinkStateEvent.unrestoredChannelCount)

	for i := 0; i < int(cLinkStateEvent.unrestoredChannelCount); i++ {
		if cUnrestoredChannels[i] != nil {
			unrestoredChannels[i] = FastSafeCGoString(cUnrestoredChannels[i])
		} else {
			unrestoredChannels[i] = ""
		}
	}

	return &LinkStateEvent{
		CurrentState:           RTM_LINK_STATE(cLinkStateEvent.currentState),
		PreviousState:          RTM_LINK_STATE(cLinkStateEvent.previousState),
		ServiceType:            RTM_SERVICE_TYPE(cLinkStateEvent.serviceType),
		Operation:              RTM_LINK_OPERATION(cLinkStateEvent.operation),
		ReasonCode:             RTM_LINK_STATE_CHANGE_REASON(cLinkStateEvent.reasonCode),
		Reason:                 FastSafeCGoString(cLinkStateEvent.reason),
		AffectedChannels:       affectedChannels,
		AffectedChannelCount:   uint(cLinkStateEvent.affectedChannelCount),
		UnrestoredChannels:     unrestoredChannels,
		UnrestoredChannelCount: uint(cLinkStateEvent.unrestoredChannelCount),
		IsResumed:              bool(cLinkStateEvent.isResumed),
		Timestamp:              uint64(cLinkStateEvent.timestamp),
	}
}

// C.struct_C_HistoryMessage to HistoryMessage
func CHistoryMessageToHistoryMessage(cHistoryMessage *C.struct_C_HistoryMessage) *HistoryMessage {
	if cHistoryMessage == nil {
		return nil
	}

	if !IsValidMemory(unsafe.Pointer(cHistoryMessage)) {
		return nil
	}

	return &HistoryMessage{
		MessageType:   RTM_MESSAGE_TYPE(cHistoryMessage.messageType),
		Message:       FastSafeCGoString(cHistoryMessage.message),
		MessageLength: uint(cHistoryMessage.messageLength),
		Timestamp:     uint64(cHistoryMessage.timestamp),
		Publisher:     FastSafeCGoString(cHistoryMessage.publisher),
		CustomType:    FastSafeCGoString(cHistoryMessage.customType),
	}
}
