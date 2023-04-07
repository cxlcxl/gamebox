<template>

  <section class="content-main-w3" id="home">
    <div class="row" style="background: #f8f8f8; margin: 0;">
      <div class="main_content col-lg-9 col-md-10 col-sm-10 col-12" style="margin: 0 auto;">
        <h2>Contact us</h2>
        <br />
        <div class="main_content">
          <p>Need help with something? Send us your question, feedback, or comment. Usually it takes us about 72 hours to react.</p>
          <p>Email: <a :href="'mailTo:'+Config.ContactEmail">{{Config.ContactEmail}}</a></p>
        </div>
      </div>
    </div>

  </section>
</template>


<script setup>
import Config from "@/setting"
</script>
