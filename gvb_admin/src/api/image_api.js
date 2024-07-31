import { Service } from "@/services/service";

export function imageRenameApi(data) {
    return Service.put("/api/images",data)
}

export function imageNameListApi(){
    return Service.get("/api/image_names")
}