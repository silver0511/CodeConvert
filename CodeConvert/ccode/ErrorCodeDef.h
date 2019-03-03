#ifndef __ERRORCODEDEF_H_
#define __ERRORCODEDEF_H_


//成功
#define RET_SUCCESS                                         0x8000000
//网络包解析失败
#define RET_ERR_NET_RECV_FAILED                             0x00000001
//系统错误偏移
#define RET_SYS_BASE                                        0x10000000
//用户错误偏移
#define RET_USER_BASE                                       0x20000000
//聊天错误偏移
#define RET_CHAT_BASE                                       0x22000000
//好友逻辑错误
#define RET_FRIEND_BASE                                     0x24000000
//群组错误
#define RET_GROUP_BASE                                      0x26000000
//公众号错误
#define RET_OFFCIAL_BASE                                    0x28000000
//商家错误
#define RET_BUSINESS_BASE                                   0x30000000

//////////////////////////////////////sys start//////////////////////////////////////
//此账号在另一个浏览器登陆了
#define RET_SYS_DISCON_USER_KICKED                          (RET_SYS_BASE + 1)
//服务器解包是失败
#define RET_UNPACK_FAILED_RESULT                            (RET_SYS_BASE + 2)
//错误的登录平台
#define RET_PLATFORM_ERROR                                  (RET_SYS_BASE + 3)
//////////////////////////////////////sys end///////////////////////////////////////

//////////////////////////////////////user start//////////////////////////////////////
//////////////////////////////////////user end///////////////////////////////////////

//////////////////////////////////////chat start//////////////////////////////////////
//////////////////////////////////////chat end///////////////////////////////////////

//////////////////////////////////////friend start//////////////////////////////////////
//////////////////////////////////////friend end///////////////////////////////////////

//////////////////////////////////////group start//////////////////////////////////////
//建群用户列表为空
#define RET_CREATE_USER_LIST_EMPTY                          (RET_GROUP_BASE + 1)
//用户操作无效
#define RET_OPERATE_TYPE_ERROR                              (RET_GROUP_BASE + 2)
//群组不存在
#define RET_GROUP_ID_INVALID                                (RET_GROUP_BASE + 3)
//用户不是群主
#define RET_GROUP_OPREATE_USER_ID_INVALID                   (RET_GROUP_BASE + 4)
//用户未加入群
#define RET_GROUP_USER_NOT_JOIN                             (RET_GROUP_BASE + 5)
//用户已经加入该群
#define RET_GROUP_USER_HAS_JOIN                             (RET_GROUP_BASE + 6)
//踢人或者邀请用户信息错误
#define RET_GROUP_OPERATE_INFO_ERROR                        (RET_GROUP_BASE + 7)
//建群超过默认最大数
#define RET_GROUP_CREATE_MAX_COUNT                          (RET_GROUP_BASE + 8)
//添加用户失败
#define RET_GROUP_INVITE_FAILED                             (RET_GROUP_BASE + 9)
//踢人失败
#define RET_GROUP_KICK_FAILED                               (RET_GROUP_BASE + 10)
//已经是群主
#define RET_GROUP_LEADER_CHANGE_SELF                        (RET_GROUP_BASE + 11)
//被转让用户参数有误
#define RET_GROUP_LEADER_NAME_IS_NIL                        (RET_GROUP_BASE + 12)
//群当前为默认加入
#define RET_GROUP_AGREE_DEFAULT                             (RET_GROUP_BASE + 13)
//群当前为需要群主同意
#define RET_GROUP_AGREE_USER                                (RET_GROUP_BASE + 14)
//////////////////////////////////////group end///////////////////////////////////////

//////////////////////////////////////offcial start//////////////////////////////////////
//////////////////////////////////////offcial end///////////////////////////////////////

//////////////////////////////////////business start//////////////////////////////////////
//////////////////////////////////////business end///////////////////////////////////////




#endif //__ERRORCODEDEF_H_