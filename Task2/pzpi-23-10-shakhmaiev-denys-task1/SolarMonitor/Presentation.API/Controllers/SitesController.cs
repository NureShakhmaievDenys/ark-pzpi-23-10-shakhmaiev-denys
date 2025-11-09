using Application.DTOs.Devices;
using Application.DTOs.Sites;
using Application.Services;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;
using System;
using System.Threading.Tasks;

namespace Presentation.API.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    [Authorize] // Весь контролер захищений
    public class SitesController : ControllerBase
    {
        private readonly SiteService _siteService;

        public SitesController(SiteService siteService)
        {
            _siteService = siteService;
        }

        private Guid GetUserIdFromToken()
        {
            var userIdString = User.FindFirstValue(ClaimTypes.NameIdentifier);
            return Guid.Parse(userIdString!);
        }

        // GET /api/sites
        [HttpGet]
        public async Task<IActionResult> GetUserSites()
        {
            var userId = GetUserIdFromToken();
            var sites = await _siteService.GetSitesForUserAsync(userId);
            return Ok(sites);
        }

        // POST /api/sites
        [HttpPost]
        public async Task<IActionResult> CreateSite([FromBody] CreateSiteRequestDto request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            var userId = GetUserIdFromToken();
            var newSite = await _siteService.CreateSiteAsync(request, userId);

            return CreatedAtAction(nameof(GetUserSites), new { id = newSite.Id }, newSite);
        }

        // POST /api/sites/{siteId}/devices
        [HttpPost("{siteId:guid}/devices")]
        public async Task<IActionResult> AddDeviceToSite([FromRoute] Guid siteId, [FromBody] CreateDeviceRequestDto request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            var userId = GetUserIdFromToken();
            var newDevice = await _siteService.AddDeviceToSiteAsync(request, siteId, userId);

            if (newDevice == null)
            {
                return NotFound(new { message = "Об'єкт не знайдено або він не належить вам." });
            }

            return StatusCode(201, newDevice);
        }
    }
}