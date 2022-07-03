export const url = `http://${window.location.hostname}:8080`;
// export const url = `http://studentsystem.xyz:8080`;  //Чтобы тестировать раскомментируй это, но когда коммитишь верни обратно либо убери с файла галочку чтобы не коммитнулось

export function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(";").shift();
}

export function logout() {
    document.cookie = "TOKEN=null;max-age=0";
    document.cookie = "ID=null;max-age=0";
    window.location.reload()
}