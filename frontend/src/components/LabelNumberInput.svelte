<script lang="ts">
    export let label: string;
    export let placeholder: string = "";
    export let value: number;
    export let min: number | null = null;
    export let max: number | null = null;
    export let valid: boolean = true;
    let customClass: string = "";
    export { customClass as class };
</script>

<div class={`flex flex-col items-start space-y-0.5
    ${customClass}`}>
    <p class="pl-1 
        text-slate-400 text-sm select-none">
        {label}
    </p>
    <input class={`w-full grow
        bg-inherit
        rounded border placeholder-slate-500 focus:outline-none
        text-l text-slate-350
        pl-2 pr-2 pt-1 pb-1
        ${!valid ? "border-red-800 focus:border-red-600" : "border-slate-500 focus:border-slate-400"}`}
        placeholder={placeholder} 
        type="number"
        min={min}
        max={max}
        bind:value={value} />
</div>