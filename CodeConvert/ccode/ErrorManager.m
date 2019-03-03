#import "ErrorManager.h"

@implementation ErrorManager
SingletonImplementation(ErrorManager)

- (id)init
{
    self = [super init];
    err_msg_dic = [[NSMutableDictionary alloc]initWithCapacity: 65535];
    [self initErrorDetail];
    return self;
}

- (void)dealloc
{
}

-(NSString*) getErrorDetail:(uint64_t) error_code
{
    NSString* err_detail = [err_msg_dic objectForKey:[NSNumber numberWithUnsignedLongLong:error_code]];
    if(err_detail == NULL)
    {
        err_detail = [NSString stringWithFormat:@"%@%lld",@"未知的错误信息:", error_code];
    }
    return err_detail;
}

-(void) setErrorDetail:(unsigned int) error_code detail:(NSString *)detail
{
    [err_msg_dic setObject:detail forKey:[NSNumber numberWithUnsignedInt:error_code]];
}

-(void) initErrorDetail
{
//////////////////////////////////////sys start//////////////////////////////////////
    [self setErrorDetail:RET_SYS_DISCON_USER_KICKED detail:@"此账号在另一个浏览器登陆了"];
    [self setErrorDetail:RET_UNPACK_FAILED_RESULT detail:@"服务器解包是失败"];
    [self setErrorDetail:RET_PLATFORM_ERROR detail:@"错误的登录平台"];
//////////////////////////////////////sys end///////////////////////////////////////

//////////////////////////////////////user start//////////////////////////////////////
//////////////////////////////////////user end///////////////////////////////////////

//////////////////////////////////////chat start//////////////////////////////////////
//////////////////////////////////////chat end///////////////////////////////////////

//////////////////////////////////////friend start//////////////////////////////////////
//////////////////////////////////////friend end///////////////////////////////////////

//////////////////////////////////////group start//////////////////////////////////////
    [self setErrorDetail:RET_CREATE_USER_LIST_EMPTY detail:@"建群用户列表为空"];
    [self setErrorDetail:RET_OPERATE_TYPE_ERROR detail:@"用户操作无效"];
    [self setErrorDetail:RET_GROUP_ID_INVALID detail:@"群组不存在"];
    [self setErrorDetail:RET_GROUP_OPREATE_USER_ID_INVALID detail:@"用户不是群主"];
    [self setErrorDetail:RET_GROUP_USER_NOT_JOIN detail:@"用户未加入群"];
    [self setErrorDetail:RET_GROUP_USER_HAS_JOIN detail:@"用户已经加入该群"];
    [self setErrorDetail:RET_GROUP_OPERATE_INFO_ERROR detail:@"踢人或者邀请用户信息错误"];
    [self setErrorDetail:RET_GROUP_CREATE_MAX_COUNT detail:@"建群超过默认最大数"];
    [self setErrorDetail:RET_GROUP_INVITE_FAILED detail:@"添加用户失败"];
    [self setErrorDetail:RET_GROUP_KICK_FAILED detail:@"踢人失败"];
    [self setErrorDetail:RET_GROUP_LEADER_CHANGE_SELF detail:@"已经是群主"];
    [self setErrorDetail:RET_GROUP_LEADER_NAME_IS_NIL detail:@"被转让用户参数有误"];
    [self setErrorDetail:RET_GROUP_AGREE_DEFAULT detail:@"群当前为默认加入"];
    [self setErrorDetail:RET_GROUP_AGREE_USER detail:@"群当前为需要群主同意"];
//////////////////////////////////////group end///////////////////////////////////////

//////////////////////////////////////offcial start//////////////////////////////////////
//////////////////////////////////////offcial end///////////////////////////////////////

//////////////////////////////////////business start//////////////////////////////////////
//////////////////////////////////////business end///////////////////////////////////////

}

@end