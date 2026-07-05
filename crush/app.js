  /* ── floating hearts ── */
  const hBg = document.getElementById('heartsBg');
  const emojis = ['💕','❤️','🌹','💖','✨','💗','🌸'];
  for (let i = 0; i < 18; i++) {
    const h = document.createElement('span');
    h.className = 'heart';
    h.textContent = emojis[Math.floor(Math.random() * emojis.length)];
    h.style.left = Math.random() * 100 + 'vw';
    h.style.fontSize = (1 + Math.random() * 1.2) + 'rem';
    const dur = 6 + Math.random() * 8;
    h.style.animationDuration = dur + 's';
    h.style.animationDelay = -(Math.random() * dur) + 's';
    hBg.appendChild(h);
  }

  /* ── dodgy No button ── */
  const noBtn = document.getElementById('noBtn');
  let noBtnW = 0, noBtnH = 0;

  // initial position: right of the Yes button
  function initNoBtn() {
    const card = document.querySelector('.card');
    const rect = card.getBoundingClientRect();
    noBtnW = noBtn.offsetWidth;
    noBtnH = noBtn.offsetHeight;
    noBtn.style.left = (rect.right - noBtnW - 24) + 'px';
    noBtn.style.top  = (rect.bottom - noBtnH - 48) + 'px';
  }
  window.addEventListener('load', initNoBtn);
  window.addEventListener('resize', initNoBtn);

  function runAway(e) {
    const mx = e.clientX !== undefined ? e.clientX : e.touches?.[0]?.clientX ?? 0;
    const my = e.clientY !== undefined ? e.clientY : e.touches?.[0]?.clientY ?? 0;

    const bx = parseFloat(noBtn.style.left) + noBtnW / 2;
    const by = parseFloat(noBtn.style.top)  + noBtnH / 2;

    const dx = bx - mx;
    const dy = by - my;
    const dist = Math.sqrt(dx * dx + dy * dy);
    const flee = 160 + Math.random() * 80;

    let nx = bx + (dx / dist) * flee - noBtnW / 2;
    let ny = by + (dy / dist) * flee - noBtnH / 2;

    // keep within viewport
    nx = Math.max(0, Math.min(window.innerWidth  - noBtnW,  nx));
    ny = Math.max(0, Math.min(window.innerHeight - noBtnH, ny));

    noBtn.style.left = nx + 'px';
    noBtn.style.top  = ny + 'px';
  }

  noBtn.addEventListener('mouseover', runAway);
  noBtn.addEventListener('touchstart', runAway, { passive: true });

  /* ── yes clicked ── */
  function sayYes() {
    document.getElementById('page1').style.display = 'none';
    const p2 = document.getElementById('page2');
    p2.classList.add('show');
    launchConfetti();
  }

  /* ── confetti ── */
  function launchConfetti() {
    const colors = ['#fff','#fde8ef','#d4a04c','#ffcce0','#ffe082','#f8bbd0'];
    for (let i = 0; i < 80; i++) {
      const c = document.createElement('div');
      c.className = 'confetti-piece';
      c.style.left  = Math.random() * 100 + 'vw';
      c.style.background = colors[Math.floor(Math.random() * colors.length)];
      c.style.width  = (6 + Math.random() * 8) + 'px';
      c.style.height = (6 + Math.random() * 8) + 'px';
      c.style.borderRadius = Math.random() > 0.5 ? '50%' : '2px';
      const dur = 2.5 + Math.random() * 2.5;
      c.style.animationDuration = dur + 's';
      c.style.animationDelay = (Math.random() * 1.2) + 's';
      document.body.appendChild(c);
      setTimeout(() => c.remove(), (dur + 1.5) * 1000);
    }
  }
