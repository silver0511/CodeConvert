package com.qbao.newim.util;
import com.qbao.newim.constdef.ErrorCodeDef;
import com.qbao.newim.qbim.R;
import java.util.HashMap;

public class ErrorDetail
{
    private static HashMap<Integer, String> error_detail_map = new HashMap<>();
    public static String GetErrorDetail(int error_code)
    {
        if(!error_detail_map.containsKey(error_code))
        {
            return "未知的错误信息:" + error_code;
        }
        return error_detail_map.get(error_code);
    }

    public static void Init()
    {
        //////////////////////////////////////sys start//////////////////////////////////////
        error_detail_map.put(ErrorCodeDef.RET_SYS_DISCON_USER_KICKED,"此账号在另一个浏览器登陆了");
        error_detail_map.put(ErrorCodeDef.RET_UNPACK_FAILED_RESULT,"服务器解包是失败");
        error_detail_map.put(ErrorCodeDef.RET_PLATFORM_ERROR,"错误的登录平台");
        //////////////////////////////////////sys end///////////////////////////////////////

        //////////////////////////////////////user start//////////////////////////////////////
        //////////////////////////////////////user end///////////////////////////////////////

        //////////////////////////////////////chat start//////////////////////////////////////
        //////////////////////////////////////chat end///////////////////////////////////////

        //////////////////////////////////////friend start//////////////////////////////////////
        //////////////////////////////////////friend end///////////////////////////////////////

        //////////////////////////////////////group start//////////////////////////////////////
        error_detail_map.put(ErrorCodeDef.RET_CREATE_USER_LIST_EMPTY,"建群用户列表为空");
        error_detail_map.put(ErrorCodeDef.RET_OPERATE_TYPE_ERROR,"用户操作无效");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_ID_INVALID,"群组不存在");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_OPREATE_USER_ID_INVALID,"用户不是群主");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_USER_NOT_JOIN,"用户未加入群");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_USER_HAS_JOIN,"用户已经加入该群");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_OPERATE_INFO_ERROR,"踢人或者邀请用户信息错误");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_CREATE_MAX_COUNT,"建群超过默认最大数");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_INVITE_FAILED,"添加用户失败");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_KICK_FAILED,"踢人失败");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_LEADER_CHANGE_SELF,"已经是群主");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_LEADER_NAME_IS_NIL,"被转让用户参数有误");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_AGREE_DEFAULT,"群当前为默认加入");
        error_detail_map.put(ErrorCodeDef.RET_GROUP_AGREE_USER,"群当前为需要群主同意");
        //////////////////////////////////////group end///////////////////////////////////////

        //////////////////////////////////////offcial start//////////////////////////////////////
        //////////////////////////////////////offcial end///////////////////////////////////////

        //////////////////////////////////////business start//////////////////////////////////////
        //////////////////////////////////////business end///////////////////////////////////////

    }
}
