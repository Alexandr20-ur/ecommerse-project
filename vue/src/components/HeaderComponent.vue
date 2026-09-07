<script setup>
    import logoSvg from '@/assets/images/logo-shape.svg'  // ← импортируем
</script>
<template>
<header class="border-bottom bg-base">
	<!-- header-main-part -->
	<section>
		<div class="container max-sm:px-2">
			<div class="d-flex max-sm:flex-wrap py-3  gap-2 align-items-center justify-content-between">
				<div class="d-flex align-items-center mr-6">
					<button class="btn btn-default btn-icon mr-2 sm:d-none">
						<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" color="currentColor" fill="none">
								<path d="M4 5L20 5" stroke="currentColor" stroke-width="2"/>
								<path d="M4 12L20 12" stroke="currentColor" stroke-width="2"/>
								<path d="M4 19L20 19" stroke="currentColor" stroke-width="2"/>
						</svg>
					</button>

					<a href="/" class="d-block flex-shrink-0 hover:opacity-80 d-flex text-decoration-none align-items-center m-0">
						<img :src="logoSvg" height="40">
						<div style="font-size:24px;" class="max-lg:hidden text-neutral-1000 text-xl font-semibold ml-1">Brand<span class="text-primary">name</span> </div>
					</a>
				</div>

				<!-- header search -->
				<div class="flex-grow max-sm:order-2 max-sm:w-full lg:mr-5">
					<form class="d-flex w-full mr-5 rounded">
						<div class="relative w-full">
							<input type="search" class="form-control bg-neutral-100 focus-none rounded-right-0" style="border:2px solid var(--primary-500); border-radius: 10px;" list="search_suggest" placeholder="Search country">
						</div>
						<button class="rounded-left-0 btn btn-primary" style="margin-left:-1px">
							Search
						</button>
					</form>
				</div>
				<!-- header search end.// -->

				<!-- header buttons -->
				<div class="d-flex gap-4 ml-10">

					<a href="#" class="d-flex group flex-col text-decoration-none text-center">
						<span class="text-base group-hover:text-primary">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"  stroke="currentColor" fill="none">
								<path d="M20.198 3H3.802C3.05147 3 2.6762 3 2.41637 3.17726C2.28768 3.26505 2.18133 3.38109 2.10567 3.51627C1.9529 3.78921 1.99024 4.15793 2.06493 4.89537C2.18958 6.12624 2.2519 6.74168 2.57823 7.18168C2.74084 7.40095 2.94701 7.58519 3.18414 7.72315C3.65999 8 4.28635 8 5.53908 8H18.4609C19.7136 8 20.34 8 20.8159 7.72315C21.053 7.58519 21.2592 7.40095 21.4218 7.18168C21.7481 6.74168 21.8104 6.12624 21.9351 4.89537C22.0098 4.15793 22.0471 3.78921 21.8943 3.51627C21.8187 3.38109 21.7123 3.26505 21.5836 3.17726C21.3238 3 20.9485 3 20.198 3Z" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
								<path d="M3 8L3 13.0408C3 16.7928 3 18.6688 4.17157 19.8344C5.34315 21 7.22876 21 11 21H13C16.7712 21 18.6569 21 19.8284 19.8344C21 18.6688 21 16.7928 21 13.0408V8" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
								<path d="M10 11H14" stroke-width="1.5" stroke-linecap="round"></path>
							</svg>
						</span>
						<span class="text-sm mt-1 text-secondary group-hover:text-underline  group-hover:text-primary">Orders</span>
					</a>

					<a href="#" class="d-flex group flex-col text-decoration-none text-center">
						<span class="text-base  group-hover:text-primary">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"  stroke="currentColor" fill="none">
								<path d="M12 20.5C12 20.5 2 14.5 2 8.69444C2 5.82563 4.10526 3.5 7 3.5C8.5 3.5 10 4 12 6C14 4 15.5 3.5 17 3.5C19.8947 3.5 22 5.82563 22 8.69444C22 14.5 12 20.5 12 20.5Z"  stroke-width="1.5" />
							</svg>
						</span>
						<span class="text-sm mt-1 text-secondary group-hover:text-underline group-hover:text-primary">Saved	</span>
					</a>

					<a href="#" class="d-flex group flex-col text-decoration-none text-center">
						<span class="text-base  group-hover:text-primary">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"  stroke="currentColor" fill="none">
							<path d="M3.87289 17.0194L2.66933 9.83981C2.48735 8.75428 2.39637 8.21152 2.68773 7.85576C2.9791 7.5 3.51461 7.5 4.58564 7.5H19.4144C20.4854 7.5 21.0209 7.5 21.3123 7.85576C21.6036 8.21152 21.5126 8.75428 21.3307 9.83981L20.1271 17.0194C19.7282 19.3991 19.5287 20.5889 18.7143 21.2945C17.9 22 16.726 22 14.3782 22H9.62182C7.27396 22 6.10003 22 5.28565 21.2945C4.47127 20.5889 4.27181 19.3991 3.87289 17.0194Z" stroke-width="1.5"></path>
							<path d="M17.5 7.5C17.5 4.46243 15.0376 2 12 2C8.96243 2 6.5 4.46243 6.5 7.5" stroke-width="1.5"></path>
							</svg>
						</span>
						<span class="text-sm mt-1 text-secondary group-hover:text-underline  group-hover:text-primary">My cart	</span>
					</a>

					<RouterLink to="/login" class="d-flex group flex-col text-decoration-none text-center">		
						<span class="text-base group-hover:text-primary">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" fill="none">
								<path d="M22 12C22 17.5228 17.5228 22 12 22C6.47715 22 2 17.5228 2 12C2 6.47715 6.47715 2 12 2C17.5228 2 22 6.47715 22 12Z" stroke="currentColor" stroke-width="1.5" />
								<path d="M15 10C15 11.6569 13.6569 13 12 13C10.3431 13 9 11.6569 9 10C9 8.34315 10.3431 7 12 7C13.6569 7 15 8.34315 15 10Z" stroke="currentColor" stroke-width="1.5" />
								<path d="M5.49994 19.0001L6.06034 18.0194C6.95055 16.4616 8.60727 15.5001 10.4016 15.5001H13.5983C15.3926 15.5001 17.0493 16.4616 17.9395 18.0194L18.4999 19.0001" stroke="currentColor" stroke-width="1.5" stroke-linecap="square" stroke-linejoin="round" />
							</svg>
						</span>
						<span class="text-sm mt-1 text-secondary group-hover:text-underline  group-hover:text-primary">Sign in </span>
					</RouterLink>
				</div>
				<!-- header buttons end.//-->
			</div>
		</div> <!-- container end.// -->
	</section>
	<!-- header-main-part end.//-->

</header> <!-- section-header end.// -->
</template>