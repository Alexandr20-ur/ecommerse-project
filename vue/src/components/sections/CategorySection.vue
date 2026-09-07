<script setup>
import { importCategory } from '@/utils/images'
const category = importCategory()
import categoryBanners from '@/assets/images/banners/category-interior.jpg'
</script>
<template>
    <section class="pb-5">
	<div class="container">
		<div class="card overflow-hidden">
			<div class="md:d-flex gap-0">
				<aside class="lg:w-70 flex-shrink-0 border-right bg-orange-100 bg-size-cover" :style="{ backgroundImage: `url(${categoryBanners})` }">
					<header class="p-6">
						<h4 class="mb-4">Home and <br> outdoor items</h4>
						<a href="#" class="btn btn-default">
							Explore all
							<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
								<path d="M9.00005 6C9.00005 6 15 10.4189 15 12C15 13.5812 9 18 9 18" stroke="#14181F" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
							</svg>
						</a>
					</header><!-- sect-heading -->
				</aside>
				<ul class="row gap-0">
					<template v-for="(cat, index) in category" :key="index">
						<li v-if="index < 10" class="col-6 lg:col-3 md:col-4  border-right border-bottom" >
						<a href="#" class="text-decoration-none p-4 text-initial d-block">
							<p class="hover:text-primary hover:text-underline">The Product name  </p>
							<img class="w-20 h-20 float-right mb-2" :src="cat">
							<p class="text-secondary"> From <br> USD 25</p>
						</a>
					</li>
					</template>
				</ul>
			</div>
			<!-- row .// -->
		</div>
	</div> <!-- container .// -->
</section>
</template>