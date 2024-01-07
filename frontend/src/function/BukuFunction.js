import axios from 'axios';

const INVENTORY_API_BASE_URL = "http://localhost:9080/api/inventory";

class BukuFunction {

    getPinjamBuku(){
        return axios.get(INVENTORY_API_BASE_URL);
    }

    createPinjamBuku(peminjam){
        return axios.post(INVENTORY_API_BASE_URL, peminjam);
    }

    getPinjamBukuById(peminjamId){
        return axios.get(INVENTORY_API_BASE_URL + '/' + peminjamId);
    }

    updatePinjamBuku(peminjam, peminjamId){
        return axios.put(INVENTORY_API_BASE_URL + '/' + peminjamId, peminjam);
    }

    deletePinjamBuku(peminjamId){
        return axios.delete(INVENTORY_API_BASE_URL + '/' + peminjamId);
    }
}

export default new BukuFunction()