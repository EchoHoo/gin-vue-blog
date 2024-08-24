import { Service } from "@/services/service";

export function createMenuApi(data){
    return Service.post('/api/menus', data)
}

export function updateMenuApi(id, data){
    return Service.put('/api/menu/' + id , data)
}

// 菜单详情
export function getMenuDetailApi(id){
    return Service.get('/api/menus/' + id)
}