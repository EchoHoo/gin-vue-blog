import { Service } from "@/services/service";

export function createAdvertApi(data) {
    return Service.post("/api/adverts", data);
}

export function updateAdvertApi(id, data) {
    return Service.put(`/api/adverts/` + id, data);
}