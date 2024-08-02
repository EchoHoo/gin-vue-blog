import { Service } from "@/services/service";

export function removeBatchApi(idList){

    return Service.delete("/api/logs",idList)
}