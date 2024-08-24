import { Service } from "@/services/service";

export function createMenuApi(data){
    return Service.post('/api/menus', data)
}

export function updateMenuApi(id, data){
    return Service.put('/api/menu/' + id , data)
}

// // 菜单详情
// export function getMenuDetailApi(id){
//     return Service.get('/api/menus/' + id)
// }

// 菜单名称列表
export function getMenuNameListApi(){
    return Service.get('/api/menu_names')
}

// 通过路径获取文章详情
export function getMenuDetailApi(path){
    return Service.get('/api/menus/detail?path=' + path)
}