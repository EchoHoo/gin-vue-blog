import { Service } from "@/services/service";

export function createMenuApi(data){
    return Service.post('/api/menus', data)
}