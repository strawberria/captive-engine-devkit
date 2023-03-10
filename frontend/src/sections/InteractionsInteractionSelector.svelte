<script lang="ts">
    import { writable, Writable } from "svelte/store";
    import LabelTextInput from "../components/LabelTextInput.svelte";
    import RadioInteraction from "../components/RadioInteraction.svelte";
    import SectionRadio from "../components/SectionRadio.svelte";
    import { selectedInteractionCriteriaIDStore, selectedInteractionIDStore, selectedInteractionResultIDStore } from "../utilities/constants";
    import { projectStore } from "../utilities/project";
    import type { ProjectInteractionData, SelectorRadioData } from "../utilities/typings";

    export let height: number | null = null;

    let filterTextStore: Writable<string> = writable("");

    let combined: string;
    let selectorRadioData: SelectorRadioData[] = [];
    function updateSelectorRadioData() {
        selectorRadioData = $projectStore.game.interactions
            .map((interactionID): [string, ProjectInteractionData] => 
                [interactionID, $projectStore.data.interactions[interactionID]])
            .filter(([interactionID, interactionData]) => {
                let actionText = interactionData.actionID !== null
                    ? $projectStore.data.actions[interactionData.actionID].name
                    : "";
                let stateText = interactionData.stateID !== null
                    ? $projectStore.data.states[interactionData.stateID].devName
                    : "";
                let componentDevNames = ["", ""];
                let componentNames = ["", ""];
                for(const index of [0, 1]) {
                    if(interactionData.componentIDs[index] === null) { continue; }
                    if(interactionData.componentIDs[index] === "anything") {
                        componentDevNames[index] = "( anything )";
                        componentNames[index] = "( anything )";
                        continue
                    }
                    const componentData = $projectStore.data[interactionData.componentTypes[0]]
                        [interactionData.componentIDs[0]];
                    if(componentData === undefined) { continue; }
                    componentNames[index] = componentData.name;
                    if((componentData as any).devName === undefined) { continue; }
                    componentDevNames[index] = (componentData as any).devName;
                }
                combined = [interactionData.devName, actionText, stateText,
                    ...componentDevNames, ...componentNames, interactionData.id]
                    .join("|").toLowerCase();

                return $filterTextStore.toLowerCase().split(",")
                    .map(term => term.trim().toLowerCase())
                    .every(term => combined.includes(term));
            })
            .map(([interactionID, interactionData]): SelectorRadioData => ({
                id: interactionID,
                component: RadioInteraction,
                props: { data: interactionData }
            }));
    }

    projectStore.subscribe(updateSelectorRadioData);
    filterTextStore.subscribe(updateSelectorRadioData);

    // Necessary to force unmount of criteria and result selectors
    const deselectable = true;
    let previousInteractionID: string | null = null;
    function handleClick(event: any) {
        const resultInteractionID = $selectedInteractionIDStore !== event.detail.id
            ? event.detail.id : deselectable
                ? null : $selectedInteractionIDStore;

        if(previousInteractionID !== resultInteractionID) {
            $selectedInteractionCriteriaIDStore = null;
            $selectedInteractionResultIDStore = null;
        }

        $selectedInteractionIDStore = resultInteractionID;
        // previousInteractionID = event.detail.id;
        previousInteractionID = resultInteractionID;
    }
</script>

<!-- allowUnlinking={false} -->
<SectionRadio height={height}
    label="Interactions"
    deleteCallback={() => { $selectedInteractionCriteriaIDStore = null; $selectedInteractionResultIDStore = null; }}
    selectedIDStore={selectedInteractionIDStore}
    selectorRadioData={selectorRadioData}
    order={$projectStore.game.interactions}
    data={$projectStore.data.interactions}
    defaultValue={{ "devName": "", "actionID": null, "componentIDs": [null, null], "stateID": null,
        "componentTypes": ["restraints", "restraints"], "order": { "criteria": [], "results": [] },
        "data": { "criteria": {}, "results": {} }}}
    on:dispatchClick={handleClick}>
    <svelte:fragment slot="pre-content">
        <div class="flex flex-row items-start w-full pr-2">
            <LabelTextInput class="w-full pb-2 mr-2"
                label="Filter Interactions (delimit terms with comma)"
                bind:value={$filterTextStore} />
        </div>
    </svelte:fragment>
</SectionRadio>