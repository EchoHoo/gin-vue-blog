import { Service } from "@/services/service";

export function createAdvertApi(data) {
    return Service.post("/api/adverts", data);
}