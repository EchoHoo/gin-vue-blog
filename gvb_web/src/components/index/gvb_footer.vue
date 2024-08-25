<template>
    <footer class="gvb_footer">
        <div class="left">
            <p class="thank">Thank For</p>
            <p class="info">

            </p>
            <p class="version">
                <span>version</span>
                <span>7.0.1</span>
            </p>
            <p>建站日期: 2024-08-11</p>

        </div>
        <div class="right">
            <div class="my_info">
                <div class="qq" title="我的QQ">
                    <!-- <a href="tencent://message/?uin=2974771769&Site=&Menu=yes" title="向我咨询">
                    </a> -->
                    <img class="show_img"
                        src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/qq_image.png"
                        alt="">
                    <img class="img"
                        src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/qq_image.png"
                        alt="">
                </div>
                <div class="wechat" title="我的微信">
                    <img class="show_img"
                        src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/wechat_image.png"
                        alt="">
                    <img class="img"
                        src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/wechat_image.png"
                        alt="">
                </div>
                <!-- <div title="我的哔哩哔哩">
                    <a href="https://space.bilibili.com/359151217" target="_blank">
                        <img class="img" src="/static/footer/bilibili_icon.svg" alt="">
                    </a>
                </div> -->
                <div title="gitee">
                    <a href="https://gitee.com/guo-cheng-hao" target="_blank">
                        <img class="img"
                            src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/logo_gitee_g_red.png"
                            alt="">
                    </a>
                </div>
                <div title="GitHub">
                    <a href="https://github.com/GuoChengH" target="_blank">
                        <img class="img"
                            src="https://gitee.com/guo-cheng-hao/drawing-bed/raw/master/%E4%B8%AA%E4%BA%BA%E5%8D%9A%E5%AE%A2/github-mark.png"
                            alt="">
                    </a>
                </div>
            </div>

        </div>
    </footer>
</template>

<script setup>

</script>

<style lang="scss">
// 底部信息栏
.gvb_footer {
    display: flex;
    padding: 25px 0;
    width: 100%;
    background-color: var(--card_bg);

    .left {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        width: 50%;
    }

    .right {
        display: flex;
        flex-direction: column;
        align-items: left;
        justify-content: center;
        width: 50%;
    }

    p {
        margin-bottom: 10px;
        font-size: 14px;

        &:last-child {
            margin-bottom: 0;
        }

        a {
            color: var(--text);
        }

    }

    .thank {
        color: #26c6da;
        font-size: 20px;
    }

    .info {
        display: flex;

        span {
            margin-right: 10px;

            &:last-child {
                margin-right: 0;
            }
        }

        img {
            width: 20px;
            margin-right: 5px;
        }
    }

    .version {
        color: white;
        display: flex;

        span {
            padding: 5px 10px;
        }

        span:first-child {
            background-color: #ff9800;
            border-bottom-left-radius: 5px;
            border-top-left-radius: 5px;
        }

        span:last-child {
            background-color: #26c6da;
            border-bottom-right-radius: 5px;
            border-top-right-radius: 5px;
        }
    }

    .beian {
        display: flex;

        img {
            width: 20px;
            margin-right: 5px;
        }
    }

    .my_info {
        display: flex;
        margin-bottom: 10px;

        >div {
            margin-right: 20px;
            cursor: pointer;

            &:last-child {
                margin-right: 0;
            }
        }

        .img {
            width: 30px;
        }

        .qq,
        .wechat {
            position: relative;

            .show_img {
                opacity: 0;
                position: absolute;
                border: 1px #f0eeee solid;
                width: 63px;
                left: -14px;
                top: -10px;
                transition: all 0.6s;
            }

            &:hover {
                .show_img {
                    opacity: 1;
                    top: -69px;
                }

            }
        }
    }

}
</style>