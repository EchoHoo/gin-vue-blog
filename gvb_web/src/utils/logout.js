import { logoutApi } from "@/api/user_api";
import { useStore } from "@/stores/store";
import { message } from "ant-design-vue";

import router from "@/router";
// 如果在js里面用router，不是从useRouter拿

export async function logout() {
    const store = useStore()
    let res = await logoutApi()
    store.clear()
    await router.push({ name: 'login' })
    if (res.code) {
        message.error(res.msg)
    }
    message.success(res.msg)

    return
}

