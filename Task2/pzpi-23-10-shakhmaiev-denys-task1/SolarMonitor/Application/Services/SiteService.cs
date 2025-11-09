using Application.DTOs.Devices;
using Application.DTOs.Sites;
using Application.Interfaces;
using Core.Entities;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Application.Services
{
    public class SiteService
    {
        private readonly IApplicationDbContext _context;

        public SiteService(IApplicationDbContext context)
        {
            _context = context;
        }

        public async Task<SiteDto> CreateSiteAsync(CreateSiteRequestDto request, Guid userId)
        {
            var site = new Site
            {
                Id = Guid.NewGuid(),
                Name = request.Name,
                Location = request.Location,
                UserId = userId
            };

            _context.Sites.Add(site);
            await _context.SaveChangesAsync();

            return new SiteDto
            {
                Id = site.Id,
                Name = site.Name,
                Location = site.Location,
                UserId = site.UserId
            };
        }

        public async Task<List<SiteDto>> GetSitesForUserAsync(Guid userId)
        {
            var sites = await _context.Sites
                .Where(s => s.UserId == userId)
                .Select(s => new SiteDto
                {
                    Id = s.Id,
                    Name = s.Name,
                    Location = s.Location,
                    UserId = s.UserId
                })
                .ToListAsync();

            return sites;
        }

        public async Task<DeviceDto?> AddDeviceToSiteAsync(CreateDeviceRequestDto request, Guid siteId, Guid userId)
        {
            var siteExists = await _context.Sites
                .AnyAsync(s => s.Id == siteId && s.UserId == userId);

            if (!siteExists)
            {
                return null;
            }

            var device = new Device
            {
                Id = Guid.NewGuid(),
                Name = request.Name,
                Brand = request.Brand,
                Config = request.Config,
                SiteId = siteId
            };

            _context.Devices.Add(device);
            await _context.SaveChangesAsync();

            return new DeviceDto
            {
                Id = device.Id,
                Name = device.Name,
                Brand = device.Brand,
                SiteId = device.SiteId
            };
        }
    }
}