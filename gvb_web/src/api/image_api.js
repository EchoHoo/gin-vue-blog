import { Service } from "@/services/service";

export function imageRenameApi(data) {
    return Service.put("/api/images",data)
}

export function imageNameListApi(){
    return Service.get("/api/image_names")
}


export function uploadImageApi(file){
    return new Promise( (resolve,reject)=>{
        const form = new FormData();
        form.append('images', file);
        Service.post("/api/images",form,{
            headers:{
                'Content-Type': 'multipart/form-data'
            }
        }) .then((res) => resolve(res))
            .catch((error) => reject(error));
})
       

}