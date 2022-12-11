import {getCookie, url} from "@/global";
import router from "@/router";

import {userData} from "@/composables/getUserData";

export function putUserData() {
    fetch(`${url}/users/${getCookie("ID")}`, {
        method: "PUT",
        headers: {
            "Authorization": `Bearer ${getCookie("TOKEN")}`
        },
        body: JSON.stringify({
            "first_name": userData.value.first_name,
            "last_name": userData.value.last_name
        })
    }).then(response => {
        if (response.ok) return response.json();
        switch (response.status) {
            case 400:
                console.log("Неверные данные");
                break;
        }
    }).then(data => {
        router.push("/auth");
        console.log(data);
    }).catch(err => {
        console.error("Cannot fetch" + err);
    });
}