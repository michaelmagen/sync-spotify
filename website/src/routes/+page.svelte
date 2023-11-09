<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { accessToken } from '$lib/auth';
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<Button class="w-40" variant="default">default</Button>
<Button variant="link" href="https://github.com/michaelmagen/sync-spotify">Link to repo</Button>
<Button variant="destructive">destructive</Button>
<Button variant="outline">outline</Button>
<Button variant="secondary">secondary</Button>
<Button variant="ghost">ghost</Button>
