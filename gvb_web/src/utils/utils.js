let timeout;

// 导出一个名为Debounce的函数，参数为func，wait（默认值为1000），immediate（默认值为false）
export function debounce(func,wait=1000,immediate=false){
    // 如果timeout不为null，则清除定时器
    if (timeout !== null) clearTimeout(timeout);
    // 如果immediate为true，则立即执行func
    if(immediate){
        let callNow = !timeout;
        // 设置定时器，wait毫秒后执行func
        timeout = setTimeout(function(){
            // 清除定时器
            timeout = null;
        },wait);
        // 如果callNow为true，则立即执行func
        if(callNow)typeof func === 'function' && func();

    }else{
        // 设置定时器，wait毫秒后执行func
        timeout = setTimeout(function(){
            // 执行func
                typeof func === 'function' && func();
        },wait)
    }
}