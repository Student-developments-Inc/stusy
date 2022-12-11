import {ref} from "vue";
import {getCookie, logout, url} from "@/global";

export let userData = ref({
    id: getCookie("ID"),
    first_name: "",
    last_name: ""
});

export function getUserData() {
    if (getCookie("ID") === undefined) logout();
    fetch(`${url}/users/${getCookie("ID")}`, {
        headers: {
            "Authorization": `Bearer ${getCookie("TOKEN")}`
        }
    }).then(response => {
        if (response.ok) return response.json();
        switch (response.status) {
            case 401:
                logout();
                break;
            case 404:
                // Сделать появление модального окна для ввода данных
                break;
        }
    }).then(data => {
        if (data !== undefined) {
            userData.value.first_name = data.first_name;
            userData.value.last_name = data.last_name;
        }
    }).catch(err => {
        console.error("Cannot fetch", err);
    });
    return {userData};
}