<script lang="ts">
	import '../app.postcss';
	import { ModeWatcher } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import { Sun, Moon, Github } from 'lucide-svelte';
	import { toggleMode } from 'mode-watcher';
	import * as Tooltip from '$lib/components/ui/tooltip';
</script>

<ModeWatcher />

<header class="w-full flex justify-center border-b">
	<div class="w-full flex justify-between py-3 px-4 max-w-7xl">
		<h1 class="scroll-m-20 text-2xl font-semibold tracking-tight">Sync Spotify</h1>
		<div class="flex justify-between items-center gap-2">
			<Tooltip.Root>
				<Tooltip.Trigger asChild let:builder>
					<Button
						builders={[builder]}
						variant="ghost"
						size="icon"
						href="https://github.com/michaelmagen/sync-spotify"
					>
						<Github />
					</Button>
				</Tooltip.Trigger>
				<Tooltip.Content>
					<p>Source Code</p>
				</Tooltip.Content>
			</Tooltip.Root>

			<Button on:click={toggleMode} variant="outline" size="icon">
				<Sun
					class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
				/>
				<Moon
					class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
				/>
				<span class="sr-only">Toggle theme</span>
			</Button>
		</div>
	</div>
</header>

<slot />
