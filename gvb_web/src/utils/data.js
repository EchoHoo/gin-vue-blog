export function getFormatDate(dateStr){
    let date = new Date(dateStr);
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let day = date.getDate();
    let hour = date.getHours();
    let minute = date.getMinutes();
    let second = date.getSeconds();
    month = (month < 10? "0" : "") + month;
    day = (day < 10? "0" : "") + day;
    hour = (hour < 10? "0" : "") + hour;
    minute = (minute < 10? "0" : "") + minute;
    second = (second < 10? "0" : "") + second;

 
    let currentDate = year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second;    
    return currentDate;
}

export function getFormatDateYMD(dateStr,joins){
    let date = new Date(dateStr);
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let day = date.getDate();
    month = (month < 10? "0" : "") + month;
    day = (day < 10? "0" : "") + day;
    let currentDate = year + "年" + month + "月" + day + "日";    
    if (joins === undefined){
        return year + "-" + month + "-" + day;
    }
    return currentDate;
}