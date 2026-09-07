<script setup>
import { importProductTech } from '@/utils/images'
const productTech = importProductTech()
</script>
<template>
    <section class="pb-5">
	<div class="container">
		<div class="card overflow-hidden">
			<div class="md:d-flex gap-x-0">
				<aside  class="lg:w-70 flex-shrink-0 max-lg:border-bottom">
					<div class="p-6">
						<header class="mb-5">
							<h4>Deals and offers</h4>
							<p class="text-secondary">Electronics & Gadgets</p>
						</header><!-- sect-heading -->
						<div class="timer mb-5">
							<div> <span class="num">04</span> <small>Days</small></div>
							<div> <span class="num">12</span> <small>Hours</small></div>
							<div> <span class="num">58</span> <small>Min</small></div>
							<div> <span class="num">02</span> <small>Sec</small></div>
						</div>
					</div>
				</aside> <!-- col.// -->
				<div class="row gap-cols-0">
					<div class="md:col sm:col-4 col-6 border-left">
						<figure class="p-2">
							<a href="#" class="d-block"> <img class="rounded" :src="productTech[1]"> </a>
							<div class="px-2 py-3 text-center">
								<p class="mb-2"><a href="#" class="text-secondary">The Product name </a></p>
								<span class="bg-red-100 text-red px-3 py-1 d-inline-block rounded-full"> -20% </span>
							</div>
						</figure>
					</div> <!-- col.// -->
					<div class="md:col sm:col-4 col-6 border-left">
						<figure class="p-2">
							<a href="#" class="d-block"> <img class="rounded" :src="productTech[2]"> </a>
							<div class="px-2 py-3 text-center">
								<p class="mb-2"><a href="#" class="text-secondary">The Product name </a></p>
								<span class="bg-red-100 text-red px-3 py-1 d-inline-block rounded-full"> -5% </span>
							</div>
						</figure>
					</div> <!-- col.// -->
					<div class="md:col sm:col-4 col-6 border-left">
						<figure class="p-2">
							<a href="#" class="d-block"> <img class="rounded" :src="productTech[3]"> </a>
							<div class="px-2 py-3 text-center">
								<p class="mb-2"><a href="#" class="text-secondary">The Product name </a></p>
								<span class="bg-red-100 text-red px-3 py-1 d-inline-block rounded-full"> -7% </span>
							</div>
						</figure>
					</div> <!-- col.// -->
					<div class="md:col sm:col-4 col-6 border-left">
						<figure class="p-2">
							<a href="#" class="d-block"> <img class="rounded" :src="productTech[4]"> </a>
							<div class="px-2 py-3 text-center">
								<p class="mb-2"><a href="#" class="text-secondary">The Product name </a></p>
								<span class="bg-red-100 text-red px-3 py-1 d-inline-block rounded-full"> -50% </span>
							</div>
						</figure>
					</div> <!-- col.// -->
					<div class="md:col sm:col-4 col-6 border-left">
						<figure class="p-2">
							<a href="#" class="d-block"> <img class="rounded" :src="productTech[5]"> </a>
							<div class="px-2 py-3 text-center">
								<p class="mb-2"><a href="#" class="text-secondary">The Product name </a></p>
								<span class="bg-red-100 text-red px-3 py-1 d-inline-block rounded-full"> -20% </span>
							</div>
						</figure>
					</div> <!-- col.// -->
				</div> <!-- row.// -->
			</div>
			<!-- row .// -->
		</div>
	</div> <!-- container .// -->
</section>
</template>