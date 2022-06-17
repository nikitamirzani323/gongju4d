<script>
    export let path_api = "";
    let listgongjuday = [];
    let listgongjunight = [];
    let tab_day = "bg-[#74aa63]"
    let tab_night = ""
    let panel_day = true
    let panel_night = false
    const handleTab = (e) => {
        switch(e){
            case "day":
                tab_day = "bg-[#74aa63]"
                tab_night = ""
                panel_day = true
                panel_night = false
                break;
            case "night":
                tab_day = ""
                tab_night = "bg-[#74aa63]"
                panel_day = false
                panel_night = true
                break;
        }
    };
    async function initgongjuDAY() {
        const resPasar = await fetch(path_api+"api/listgongjuday", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({}),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
                let record = jsonPasar.record;
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listgongjuday = [
                            ...listgongjuday,
                            {
                                sdsbday_date: record[i]["sdsbday_date"],
                                sdsbday_prize1: record[i]["sdsbday_prize1"],
                                sdsbday_prize2: record[i]["sdsbday_prize2"],
                                sdsbday_prize3: record[i]["sdsbday_prize3"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            } else {
                alert("Error");
            }
        }
    }
    // initgongjuDAY()
</script>
<section class="mt-5 mb-5 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-4 left-10 z-auto">결과</h2>
</section>
<ul class="flex justify-center w-full gap-5">
    <li on:click={() => {
        handleTab("day");
    }} class="border-2 border-[#74aa63] {tab_day}  py-2 px-5 rounded-lg cursor-pointer">공주의 날</li>
    <li on:click={() => {
        handleTab("night");
    }} class="border-2 border-[#74aa63] {tab_night} py-2 px-5 rounded-lg cursor-pointer">공주의 밤</li>
</ul>
<div class="overflow-x-auto mt-5 mb-10">
    {#if panel_day}
        <table class="table w-full">
            <thead >
                <tr>
                <th class="bg-white border-2 border-[pink] text-center">데이트</th>
                <th class="bg-white border-2 border-[pink] text-center">큰 상</th>
                <th class="bg-white border-2 border-[pink] text-center">2등상</th>
                <th class="bg-white border-2 border-[pink] text-center">3등상</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
            </tbody>
        </table>
    {/if}
    {#if panel_night}
        <table class="table w-full">
            <thead >
                <tr>
                <th class="bg-white border-2 border-[pink] text-center">데이트</th>
                <th class="bg-white border-2 border-[pink] text-center">큰 상</th>
                <th class="bg-white border-2 border-[pink] text-center">2등상</th>
                <th class="bg-white border-2 border-[pink] text-center">3등상</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">March 4, 2022</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">2134</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">1532</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">9578</td>
                </tr>
            </tbody>
        </table>
    {/if}
</div>