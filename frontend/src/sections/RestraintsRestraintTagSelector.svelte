<script lang="ts">
    import { writable, Writable } from "svelte/store";
    import IconButton from "../components/IconButton.svelte";
    import LabelTextInput from "../components/LabelTextInput.svelte";
    import RadioRestraintTag from "../components/RadioRestraintTag.svelte";
    import SectionRadio from "../components/SectionRadio.svelte";
    import { selectedRestraintIDStore } from "../utilities/constants";
    import { projectStore } from "../utilities/project";
    import type { SelectorRadioData } from "../utilities/typings";

    export let height: number | null = null;

    let selectorRadioData: SelectorRadioData[] = [];
    // Somehow subscribe not working
    function updateSelectorRadioData() {
        if($selectedRestraintIDStore === null) { return; }
        selectorRadioData = Object.entries($projectStore.data.restraints[$selectedRestraintIDStore].tags)
            .map(([index, tag]): SelectorRadioData => ({
                id: index,
                component: RadioRestraintTag,
                props: { index: index, tag: tag }
            })); 
    }
    projectStore.subscribe(updateSelectorRadioData);
    selectedRestraintIDStore.subscribe(updateSelectorRadioData);

    let dummyIDStore: Writable<string | null> = writable(null)
    let restraintTagInput: string = "";
    function submitAddTag() {
        if(restraintTagInput === ""
            || $projectStore.data.restraints[$selectedRestraintIDStore].tags
                .indexOf(restraintTagInput) !== -1) { return; }
        $projectStore.data.restraints[$selectedRestraintIDStore].tags.push(restraintTagInput);
        $projectStore = $projectStore;
        restraintTagInput = "";
    }
    function submitDeleteTag(event) {
        const index = event.detail.index;
        $projectStore.data.restraints[$selectedRestraintIDStore].tags.splice(index, 1);
        $projectStore = $projectStore;
    }
</script>

<SectionRadio height={height} 
    label="Restraint Tags"
    nobuttons={true}
    noheader={true}
    overridepointer={true}
    selectedIDStore={dummyIDStore}
    selectorRadioData={selectorRadioData}
    order={[]}
    data={{}}
    defaultValue={null}
    on:dispatchClick={submitDeleteTag}>
    <svelte:fragment slot="pre-content">
        <div class="flex flex-row w-full space-x-3 pb-2">
            <LabelTextInput class="grow"
                bind:value={restraintTagInput}
                placeholder={"locked-star_key"}
                on:submit={submitAddTag} />
            <IconButton label="Add"
                onclick={submitAddTag} />
        </div>
    </svelte:fragment>
</SectionRadio>