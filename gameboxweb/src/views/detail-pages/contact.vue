<template>

  <section class="content-main-w3" id="home">
    <div class="page-width" style="background: #f8f8f8; margin: 0 auto; padding: 20px;">
      <h2>Contact us</h2>
      <br />
      <div class="main_content">
        <p>Need help with something? Send us your question, feedback, or comment. Usually it takes us about 72 hours to react.</p>
        <p>Email: <a :href="'mailTo:'+Config.ContactEmail">{{Config.ContactEmail}}</a></p>
      </div>
    </div>

  </section>
</template>


<script setup>
import Config from "@/setting"
</script>
