package com.qbao.newim.constdef;


public interface PackTypeDef
{
    short

    //客户端包起始地址
    NEW_DEF_DEFAULT_CLIENT_PACKET_BASE                  = (0),
    //系统级消息起始地址
    NEW_DEF_SYS_BASE                                    = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 0),
    //用户消息起始地址
    NEW_DEF_USER_BASE                                   = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE +200),
    //聊天消息起始地址
    NEW_DEF_CHAT_BASE                                   = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 3000),
    //好友消息起始地址
    NEW_DEF_FRIEND_BASE                                 = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 4000),
    //群消息起始地址
    NEW_DEF_GROUP_BASE                                  = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 5000),
    //公众号起始地址
    NEW_DEF_OFFCIAL_BASE                                = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 6000),
    //商家小二起始地址
    NEW_DEF_BUSINESS_BASE                               = (NEW_DEF_DEFAULT_CLIENT_PACKET_BASE + 6500),

//////////////////////////////////////sys start//////////////////////////////////////
    //握手包
    NEW_DEF_HANDSHAKE_RQ                                = (NEW_DEF_SYS_BASE + 1),
    NEW_DEF_HANDSHAKE_RS                                = (NEW_DEF_SYS_BASE + 2),
    //握手ack
    NEW_DEF_ACK_ID                                      = (NEW_DEF_SYS_BASE + 3),
    //心跳包
    NEW_DEF_HEART_RQ                                    = (NEW_DEF_SYS_BASE + 4),
    NEW_DEF_HEART_RS                                    = (NEW_DEF_SYS_BASE + 5),
    //服务器主动断开连接通知
    NEW_DEF_SERVER_DISCON_ID                            = (NEW_DEF_SYS_BASE + 6),
    //错误包的type
    NEW_DEF_SERVER_BASE_ERROR                           = (NEW_DEF_SYS_BASE + 10),
    //登录
    NEW_DEF_LOGIN_RQ                                    = (NEW_DEF_SYS_BASE + 20),
    NEW_DEF_LOGIN_RS                                    = (NEW_DEF_SYS_BASE + 21),
    //更新系统时间
    NEW_UPDATE_TIME_ID                                  = (NEW_DEF_SYS_BASE + 22),
//////////////////////////////////////sys end///////////////////////////////////////

//////////////////////////////////////user start//////////////////////////////////////
    //搜索用户
    NEW_DEF_USER_SEARCH_RQ                              = (NEW_DEF_USER_BASE + 1),
    NEW_DEF_USER_SEARCH_RS                              = (NEW_DEF_USER_BASE + 2),
    //修改用户信息
    NEW_DEF_USER_CHANGE_RQ                              = (NEW_DEF_USER_BASE + 5),
    NEW_DEF_USER_CHANGE_RS                              = (NEW_DEF_USER_BASE + 6),
    //获取用户信息
    NEW_DEF_USER_INFO_RQ                                = (NEW_DEF_USER_BASE + 9),
    NEW_DEF_USER_INFO_RS                                = (NEW_DEF_USER_BASE + 10),
    //属性变化包
    NEW_DEF_USER_ATTR_ID                                = (NEW_DEF_USER_BASE + 13),
//////////////////////////////////////user end///////////////////////////////////////

//////////////////////////////////////chat start//////////////////////////////////////
    //消息发送
    NEW_DEF_CHAT_CLIENT_SEND_MESSAGE_RQ                 = (NEW_DEF_CHAT_BASE + 1),
    NEW_DEF_CHAT_CLIENT_SEND_MESSAGE_RS                 = (NEW_DEF_CHAT_BASE + 2),
    NEW_DEF_CHAT_SERVER_SEND_MESSAGE_RQ                 = (NEW_DEF_CHAT_BASE + 3),
    NEW_DEF_CHAT_SERVER_SEND_MESSAGE_RS                 = (NEW_DEF_CHAT_BASE + 4),
    //离线消息获取
    NEW_DEF_CHAT_GET_OFFLINE_MESSAGE_RQ                 = (NEW_DEF_CHAT_BASE + 5),
    NEW_DEF_CHAT_GET_OFFLINE_MESSAGE_RS                 = (NEW_DEF_CHAT_BASE + 6),
//////////////////////////////////////chat end///////////////////////////////////////

//////////////////////////////////////friend start//////////////////////////////////////
    //获取好友列表
    NEW_DEF_FRIEND_LIST_RQ                              = (NEW_DEF_FRIEND_BASE + 1),
    NEW_DEF_FRIEND_LIST_RS                              = (NEW_DEF_FRIEND_BASE + 2),
    //添加好友
    NEW_DEF_CLIENT_FRIEND_ADD_RQ                        = (NEW_DEF_FRIEND_BASE + 5),
    NEW_DEF_CLIENT_FRIEND_ADD_RS                        = (NEW_DEF_FRIEND_BASE + 6),
    NEW_DEF_SERVER_FRIEND_ADD_RQ                        = (NEW_DEF_FRIEND_BASE + 7),
    NEW_DEF_SERVER_FRIEND_ADD_RS                        = (NEW_DEF_FRIEND_BASE + 8),
    //好友确认
    NEW_DEF_SERVER_FRIEND_CONFIRM_RQ                    = (NEW_DEF_FRIEND_BASE + 9),
    NEW_DEF_SERVER_FRIEND_CONFIRM_RS                    = (NEW_DEF_FRIEND_BASE + 10),
    //删除好友操作
    NEW_DEF_FRIEND_DEL_RQ                               = (NEW_DEF_FRIEND_BASE + 11),
    NEW_DEF_FRIEND_DEL_RS                               = (NEW_DEF_FRIEND_BASE + 12),
    //修改备注
    NEW_DEF_FRIEND_REMARK_RQ                            = (NEW_DEF_FRIEND_BASE + 13),
    NEW_DEF_FRIEND_REMARK_RS                            = (NEW_DEF_FRIEND_BASE + 14),
    NEW_DEF_CLIENT_FRIEND_CONFIRM_RQ                    = (NEW_DEF_FRIEND_BASE + 15),
    NEW_DEF_CLINET_FRIEND_CONFIRM_RS                    = (NEW_DEF_FRIEND_BASE + 16),
//////////////////////////////////////friend end///////////////////////////////////////

//////////////////////////////////////group start//////////////////////////////////////
    //创建群
    NEW_DEF_GROUP_CREATE_RQ                             = (NEW_DEF_GROUP_BASE + 1),
    NEW_DEF_GROUP_CREATE_RS                             = (NEW_DEF_GROUP_BASE + 2),
    //群成员操作(申请，同意，拒绝，邀请)
    NEW_DEF_GROUP_OP_RQ                                 = (NEW_DEF_GROUP_BASE + 7),
    NEW_DEF_GROUP_OP_RS                                 = (NEW_DEF_GROUP_BASE + 8),
    //获取群信息
    NEW_DEF_GROUP_DETAIL_INFO_RQ                        = (NEW_DEF_GROUP_BASE + 11),
    NEW_DEF_GROUP_DETAIL_INFO_RS                        = (NEW_DEF_GROUP_BASE + 12),
    //获取群列表
    NEW_DEF_GROUP_LIST_RQ                               = (NEW_DEF_GROUP_BASE + 15),
    NEW_DEF_GROUP_LIST_RS                               = (NEW_DEF_GROUP_BASE + 16),
    //群修改设置
    NEW_DEF_GROUP_CHANGE_RQ                             = (NEW_DEF_GROUP_BASE + 22),
    NEW_DEF_GROUP_CHANGE_RS                             = (NEW_DEF_GROUP_BASE + 23),
    //群消息
    NEW_DEF_GROUP_CLIENT_SEND_MESSAGE_RQ                = (NEW_DEF_GROUP_BASE + 24),
    NEW_DEF_GROUP_CLIENT_SEND_MESSAGE_RS                = (NEW_DEF_GROUP_BASE + 25),
    NEW_DEF_GROUP_SERVER_SEND_MESSAGE_RQ                = (NEW_DEF_GROUP_BASE + 26),
    NEW_DEF_GROUP_SERVER_SEND_MESSAGE_RS                = (NEW_DEF_GROUP_BASE + 27),
    //群离线消息
    NEW_DEF_GROUP_GET_OFFLINE_MESSAGE_RQ                = (NEW_DEF_GROUP_BASE + 28),
    NEW_DEF_GROUP_GET_OFFLINE_MESSAGE_RS                = (NEW_DEF_GROUP_BASE + 29),
    //获取校验信息
    NEW_DEF_GROUP_VERIFY_MESSAGE_RQ                     = (NEW_DEF_GROUP_BASE + 30),
    NEW_DEF_GROUP_VERIFY_MESSAGE_RS                     = (NEW_DEF_GROUP_BASE + 31),
    //获取群聊漏掉的消息
    NEW_DEF_GROUP_MSG_LEAVE_OUT_RQ                      = (NEW_DEF_GROUP_BASE + 32),
    NEW_DEF_GROUP_MSG_LEAVE_OUT_RS                      = (NEW_DEF_GROUP_BASE + 33),
    //修改群成员
    NEW_DEF_GROUP_MODIFY_RQ                             = (NEW_DEF_GROUP_BASE + 34),
    NEW_DEF_GROUP_MODIFY_RS                             = (NEW_DEF_GROUP_BASE + 35),
    //修改群成员 服务器主动推送
    NEW_DEF_GROUP_MODIFY_SERVER_RQ                      = (NEW_DEF_GROUP_BASE + 36),
    NEW_DEF_GROUP_MODIFY_SERVER_RS                      = (NEW_DEF_GROUP_BASE + 37),
    //创建群 服务器主动推送
    NEW_DEF_GROUP_CREATE_SERVER_RQ                      = (NEW_DEF_GROUP_BASE + 38),
    NEW_DEF_GROUP_CREATE_SERVER_RS                      = (NEW_DEF_GROUP_BASE + 39),
    //群主转让
    NEW_DEF_GROUP_LEADER_CHANGE_RQ                      = (NEW_DEF_GROUP_BASE + 40),
    NEW_DEF_GROUP_LEADER_CHANGE_RS                      = (NEW_DEF_GROUP_BASE + 41),
//////////////////////////////////////group end///////////////////////////////////////

//////////////////////////////////////offcial start//////////////////////////////////////
    //公众号消息发送
    NEW_DEF_OFFCIAL_CLIENT_MESSAGE_RQ                   = (NEW_DEF_OFFCIAL_BASE + 1),
    NEW_DEF_OFFCIAL_CLIENT_MESSAGE_RS                   = (NEW_DEF_OFFCIAL_BASE + 2),
    NEW_DEF_OFFCIAL_SERVER_MESSAGE_RQ                   = (NEW_DEF_OFFCIAL_BASE + 3),
    NEW_DEF_OFFCIAL_SERVER_MESSAGE_RS                   = (NEW_DEF_OFFCIAL_BASE + 4),
    //公众号离线消息
    NEW_DEF_OFFCIAL_GET_OFFLINE_MESSAGE_RQ              = (NEW_DEF_OFFCIAL_BASE + 5),
    NEW_DEF_OFFCIAL_GET_OFFLINE_MESSAGE_RS              = (NEW_DEF_OFFCIAL_BASE + 6),
//////////////////////////////////////offcial end///////////////////////////////////////

//////////////////////////////////////business start//////////////////////////////////////
//////////////////////////////////////business end///////////////////////////////////////

    END_OF_TYPE = 19999;
}
