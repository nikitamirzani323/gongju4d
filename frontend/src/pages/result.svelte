<script>
    export let path_api = "";
    let listgongjuday = [];
    let listgongjunight = [];
    let tab_day = "bg-[#74aa63]"
    let tab_night = ""
    
    listgongjuday = [
        {tgl:"March 1, 2022",prize_1:"1134",prize_2:"1532",prize_3:"0923"},
        {tgl:"March 2, 2022",prize_1:"2134",prize_2:"2532",prize_3:"1923"},
        {tgl:"March 3, 2022",prize_1:"3134",prize_2:"3532",prize_3:"2923"},
        {tgl:"March 4, 2022",prize_1:"4134",prize_2:"4532",prize_3:"3923"},
        {tgl:"March 5, 2022",prize_1:"5134",prize_2:"5532",prize_3:"4923"},
        {tgl:"March 6, 2022",prize_1:"6134",prize_2:"6532",prize_3:"5923"},
        {tgl:"March 7, 2022",prize_1:"7134",prize_2:"7532",prize_3:"6923"},
        {tgl:"March 8, 2022",prize_1:"8134",prize_2:"8532",prize_3:"7923"},
        {tgl:"March 9, 2022",prize_1:"9134",prize_2:"9532",prize_3:"8923"},
        {tgl:"March 10, 2022",prize_1:"0134",prize_2:"0532",prize_3:"9923"},
        {tgl:"March 11, 2022",prize_1:"2234",prize_2:"1632",prize_3:"0023"},
        {tgl:"March 12, 2022",prize_1:"2334",prize_2:"1732",prize_3:"0123"},
        {tgl:"March 13, 2022",prize_1:"2434",prize_2:"1832",prize_3:"0223"},
        {tgl:"March 14, 2022",prize_1:"2534",prize_2:"1932",prize_3:"0323"},
        {tgl:"March 15, 2022",prize_1:"2634",prize_2:"1032",prize_3:"0423"},
        {tgl:"March 16, 2022",prize_1:"2734",prize_2:"1132",prize_3:"0523"},
    ];

    const handleTab = (e) => {
        switch(e){
            case "day":
                tab_day = "bg-[#74aa63]"
                tab_night = ""
                listgongjuday = [
                    {tgl:"March 1, 2022",prize_1:"1134",prize_2:"1532",prize_3:"0923"},
                    {tgl:"March 2, 2022",prize_1:"2134",prize_2:"2532",prize_3:"1923"},
                    {tgl:"March 3, 2022",prize_1:"3134",prize_2:"3532",prize_3:"2923"},
                    {tgl:"March 4, 2022",prize_1:"4134",prize_2:"4532",prize_3:"3923"},
                    {tgl:"March 5, 2022",prize_1:"5134",prize_2:"5532",prize_3:"4923"},
                    {tgl:"March 6, 2022",prize_1:"6134",prize_2:"6532",prize_3:"5923"},
                    {tgl:"March 7, 2022",prize_1:"7134",prize_2:"7532",prize_3:"6923"},
                    {tgl:"March 8, 2022",prize_1:"8134",prize_2:"8532",prize_3:"7923"},
                    {tgl:"March 9, 2022",prize_1:"9134",prize_2:"9532",prize_3:"8923"},
                    {tgl:"March 10, 2022",prize_1:"0134",prize_2:"0532",prize_3:"9923"},
                    {tgl:"March 11, 2022",prize_1:"2234",prize_2:"1632",prize_3:"0023"},
                    {tgl:"March 12, 2022",prize_1:"2334",prize_2:"1732",prize_3:"0123"},
                    {tgl:"March 13, 2022",prize_1:"2434",prize_2:"1832",prize_3:"0223"},
                    {tgl:"March 14, 2022",prize_1:"2534",prize_2:"1932",prize_3:"0323"},
                    {tgl:"March 15, 2022",prize_1:"2634",prize_2:"1032",prize_3:"0423"},
                    {tgl:"March 16, 2022",prize_1:"2734",prize_2:"1132",prize_3:"0523"},
                ];
                break;
            case "night":
                tab_day = ""
                tab_night = "bg-[#74aa63]"
                listgongjuday = [
                    {tgl:"March 1, 2022",prize_1:"1134",prize_2:"1532",prize_3:"0923"},
                    {tgl:"March 2, 2022",prize_1:"2144",prize_2:"2532",prize_3:"1923"},
                    {tgl:"March 3, 2022",prize_1:"3154",prize_2:"3532",prize_3:"2923"},
                    {tgl:"March 4, 2022",prize_1:"4164",prize_2:"4532",prize_3:"3923"},
                    {tgl:"March 5, 2022",prize_1:"5174",prize_2:"5532",prize_3:"4923"},
                    {tgl:"March 6, 2022",prize_1:"6184",prize_2:"6532",prize_3:"5923"},
                    {tgl:"March 7, 2022",prize_1:"7194",prize_2:"7532",prize_3:"6923"},
                    {tgl:"March 8, 2022",prize_1:"8104",prize_2:"8532",prize_3:"7923"},
                    {tgl:"March 9, 2022",prize_1:"9115",prize_2:"9532",prize_3:"8923"},
                    {tgl:"March 10, 2022",prize_1:"0136",prize_2:"0532",prize_3:"9923"},
                    {tgl:"March 11, 2022",prize_1:"2237",prize_2:"1632",prize_3:"0023"},
                    {tgl:"March 12, 2022",prize_1:"2338",prize_2:"1732",prize_3:"0123"},
                    {tgl:"March 13, 2022",prize_1:"2439",prize_2:"1832",prize_3:"0223"},
                    {tgl:"March 14, 2022",prize_1:"2530",prize_2:"1932",prize_3:"0323"},
                    {tgl:"March 15, 2022",prize_1:"2631",prize_2:"1032",prize_3:"0423"},
                    {tgl:"March 16, 2022",prize_1:"2732",prize_2:"1132",prize_3:"0523"},
                ];
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
    <h2 class="text-[pink] text-lg lg:text-3xl text-center bg-white absolute -top-4 left-10 z-auto">결과</h2>
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
    <table class="table w-full">
        <thead class="sticky top-0">
            <tr>
                <th class="bg-white border-2 border-[pink] text-center text-xs lg:text-lg">데이트</th>
                <th class="bg-white border-2 border-[pink] text-center text-xs lg:text-lg">1등상</th>
                <th class="bg-white border-2 border-[pink] text-center text-xs lg:text-lg">2등상</th>
                <th class="bg-white border-2 border-[pink] text-center text-xs lg:text-lg">3등상</th>
            </tr>
        </thead>
        <tbody>
            {#each listgongjuday as rec}
                <tr>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">{rec.tgl}</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">{rec.prize_1}</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">{rec.prize_2}</td>
                    <td class="bg-white border-2 border-[pink] text-[11px] text-sm text-center">{rec.prize_3}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>